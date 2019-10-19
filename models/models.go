package models

import (
	"time"

	"github.com/idirall22/post/models"
)

// Group structure
type Group struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	AdminID    int64      `json:"admin_id"`
	TimelineID int64      `json:"timeline_id"`
	UsersIDs   []int64    `json:"users_ids"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

// Timeline structure
type Timeline struct {
	ID        int64         `json:"id"`
	GroupID   int64         `json:"group_id"`
	Posts     []models.Post `json:"posts"`
	CreatedAt time.Time     `json:"created_at"`
	DeletedAt *time.Time    `json:"deleted_at"`
}
