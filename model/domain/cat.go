package domain

import (
	"time"
)

type Cat struct {
	Id          int
	UserId      int
	Name        string
	Race        string
	Sex         string
	AgeInMonth  int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
