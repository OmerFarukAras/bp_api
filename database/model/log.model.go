package model

import (
	"time"
)

type Log struct {
	ID        string    `storm:"index"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
