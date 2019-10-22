package models

import (
	"time"

	"github.com/idirall22/post/models"
)

// Group structure
type Group struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	AdminID   int64         `json:"admin_id"`
	UsersIDs  []int64       `json:"users_ids"`
	Posts     []models.Post `json:"posts"`
	CreatedAt time.Time     `json:"created_at"`
	DeletedAt *time.Time    `json:"deleted_at"`
}
