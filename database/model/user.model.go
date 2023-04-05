package model

import (
	"time"
)

type User struct {
	ID string `storm:"index"`

	Username string `json:"content"`

	PostID []string `json:"posts"`

	Details    string `json:"details"`
	ProfilePic string `json:"profilePic"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedA"`
}
