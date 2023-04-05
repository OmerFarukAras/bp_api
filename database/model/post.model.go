package model

import (
	"time"
)

type Post struct {
	ID        string    `storm:"index"`
	Title     string    `json:"content"`
	Content   string    `json:"details"`
	Author    string    `json:"author"`
	MetaTags  []string  `json:"MetaTags"`
	Views     int       `json:"int"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedA"`
}
