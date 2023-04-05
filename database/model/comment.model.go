package model

import (
	"time"
)

type Comment struct {
	ID        string    `storm:"index"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Comments  []Comment `json:"comments"`
	Likes     []string  `json:"likes"`
	CommentOf string    `json:"CommentOf"`
}
