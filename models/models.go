package models

import "time"

// Group structure
type Group struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	AdminID   int64      `json:"admin_id"`
	UsersIDs  []int64    `json:"users_ids"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
