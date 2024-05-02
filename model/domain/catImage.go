package domain

import (
	"time"
)

type CatImage struct {
	Id        int
	CatID     int
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
