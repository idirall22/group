package group

import (
	"context"

	"github.com/idirall22/group/models"
	"github.com/idirall22/utilities"
)

// Add a group
func (s *Service) addGroup(ctx context.Context, form GForm) (int64, error) {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return 0, err
	}

	if err := form.ValidateForm(); err != nil {
		return 0, err
	}

	return s.provider.New(ctx, userID, form.Name)
}

// Get a group
func (s *Service) getGroup(ctx context.Context, id int64) (*models.Group, error) {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	if err := utilities.ValidateID(id); err != nil {
		return nil, err
	}

	return s.provider.Get(ctx, id, userID)
}

func (s *Service) listGroups(ctx context.Context, limit, offset int) ([]*models.Group, error) {

	off, lim := utilities.ValidateLimitAndOffset(offset, limit, DefaultGroupLimit, MaxGroupLimit)

	return s.provider.List(ctx, off, lim)
}

// Update a group
func (s *Service) updateGroup(ctx context.Context, id int64, form GForm) error {

	adminID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return err
	}

	if err := utilities.ValidateID(id); err != nil {
		return err
	}

	if err := form.ValidateForm(); err != nil {
		return err
	}

	return s.provider.Update(ctx, id, adminID, form.Name)
}

// Delete a group
func (s *Service) deleteGroup(ctx context.Context, id int64) error {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return err
	}

	if err := utilities.ValidateID(id); err != nil {
		return err
	}

	return s.provider.Delete(ctx, id, userID)
}

// Join a Group
func (s *Service) joinGroup(ctx context.Context, id int64) error {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return err
	}

	if err := utilities.ValidateID(id); err != nil {
		return err
	}

	return s.provider.Join(ctx, userID, id)
}

// Leave a group
func (s *Service) leaveGroup(ctx context.Context, id int64) error {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return err
	}

	if err := utilities.ValidateID(id); err != nil {
		return err
	}

	return s.provider.Leave(ctx, userID, id)
}
