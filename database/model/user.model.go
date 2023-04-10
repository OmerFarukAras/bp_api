package model

import (
	"bp_api/util"
	"time"
)

type User struct {
	ID string `storm:"index"`

	Username string `json:"content"`

	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`

	PostID []string `json:"posts"`

	Details    string `json:"details"`
	ProfilePic string `json:"profilePic"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (U *User) CreateID() {
	U.ID = util.CreateUID()
}

func (U *User) GetName() string {
	return U.FirstName + " " + U.LastName
}

func (U *User) SetCreationTime() {
	U.CreatedAt = time.Now()
}

func (U *User) SetUpdateTime() {
	U.UpdatedAt = time.Now()
}
