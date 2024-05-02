package domain

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
