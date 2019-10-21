package provider

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/idirall22/group/models"
	postModels "github.com/idirall22/post/models"
	"github.com/lib/pq"
)

// PostgresProvider structure
type PostgresProvider struct {
	DB        *sql.DB
	TableName string
}

// New add a group
func (p *PostgresProvider) New(ctx context.Context, userID int64, name string) (int64, error) {

	tx, err := p.DB.BeginTx(ctx, nil)

	defer tx.Rollback()

	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM groups WHERE name='%s')`, name)

	stmt, err := tx.PrepareContext(ctx, query)

	if err != nil {
		return 0, err
	}

	exists := false

	if err := stmt.QueryRowContext(ctx).Scan(&exists); err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("Group with the same name already exists")
	}

	query = fmt.Sprintf(`INSERT INTO %s (name, admin_id) VALUES ($1, $2) RETURNING id`,
		p.TableName)

	stmt, err = tx.PrepareContext(ctx, query)

	if err != nil {
		return 0, nil
	}

	groupID := int64(0)

	if err := stmt.QueryRowContext(ctx, name, userID).Scan(&groupID); err != nil {
		return 0, err
	}

	tx.Commit()

	return groupID, nil
}

// Get get a group
func (p *PostgresProvider) Get(ctx context.Context, id, userID int64, name string) (*models.Group, error) {

	tx, err := p.DB.BeginTx(ctx, nil)

	defer tx.Rollback()

	if err != nil {
		return nil, err
	}

	// Check if the user has already joined the group
	query := fmt.Sprintf(`
		SELECT EXISTS (SELECT 1 FROM %s WHERE id=%d AND admin_id=%d OR %d=ANY (users_ids))`,
		p.TableName, id, userID, userID)

	stmt, err := tx.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	exists := false

	if err := stmt.QueryRowContext(ctx).Scan(&exists); err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrorMustJoinGroup
	}

	// query by id if id > 0 or by name
	query = fmt.Sprintf(`
		SELECT
		g.name, g.admin_id, g.users_ids,
		p.id, p.content, p.media_urls, p.user_id, p.created_at
		FROM
		(SELECT name, admin_id, users_ids FROM %s WHERE id=%d) g
		JOIN
		posts p
		ON p.group_id=%d
	`, p.TableName, id, id)

	stmt, err = tx.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	group := &models.Group{}

	for rows.Next() {
		post := postModels.Post{}

		// fmt.Println(rows.Columns())
		if err := rows.Scan(
			&group.Name,
			&group.AdminID,
			pq.Array(&group.UsersIDs),
			&post.ID,
			&post.Content,
			pq.Array(&post.MediaURLs),
			&post.UserID,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}
		group.Posts = append(group.Posts, post)
	}

	group.ID = id

	tx.Commit()

	return group, nil
}

// List get a list of groups
func (p *PostgresProvider) List(ctx context.Context, offset, limit int) ([]*models.Group, error) {
	query := fmt.Sprintf(`SELECT id, name, admin_id, created_at FROM %s LIMIT %d OFFSET %d`,
		p.TableName, offset, limit*offset)

	stmt, err := p.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	groups := []*models.Group{}

	for rows.Next() {
		group := &models.Group{}
		if err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.AdminID,
			&group.CreatedAt,
		); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// Update update a group
func (p *PostgresProvider) Update(ctx context.Context, id int64, name string) error {

	query := fmt.Sprintf(`UPDATE %s SET name='%s' WHERE id=%d`, p.TableName, name, id)

	_, err := p.DB.ExecContext(ctx, query)

	return err
}

// Delete delete a group
func (p *PostgresProvider) Delete(ctx context.Context, id, userID int64) error {

	query := fmt.Sprintf(`UPDATE %s SET deleted_at=now() WHERE id=%d AND admin_id=%d`,
		p.TableName, id, userID)

	stmt, err := p.DB.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx)
	return err
}

// Join join a group
func (p *PostgresProvider) Join(ctx context.Context, userID, groupID int64) error {

	tx, err := p.DB.BeginTx(ctx, nil)

	defer tx.Rollback()

	if err != nil {
		return err
	}

	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM users WHERE id=%d)`, userID)

	stmt, err := tx.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	exists := false
	err = stmt.QueryRowContext(ctx).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("User not exists")
	}

	query = fmt.Sprintf(`UPDATE %s SET users_ids=array_append(users_ids, %d) WHERE id=%d AND admin_id != %d`,
		p.TableName, userID, groupID, userID)

	stmt, err = tx.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx)

	tx.Commit()

	return err
}

// Leave leave a group
func (p *PostgresProvider) Leave(ctx context.Context, userID, groupID int64) error {

	tx, err := p.DB.BeginTx(ctx, nil)

	defer tx.Rollback()

	if err != nil {
		return err
	}

	// Check if group exists
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE id=%d)`, p.TableName, groupID)

	stmt, err := tx.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	exists := false

	err = stmt.QueryRowContext(ctx).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return ErrorGroupNotExists
	}

	// Remove a user from users_ids array
	query = fmt.Sprintf(`UPDATE %s SET users_ids=array_remove(users_ids, %d) WHERE id=%d`,
		p.TableName, userID, groupID)

	stmt, err = tx.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
