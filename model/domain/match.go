package domain

import (
	"time"
)

type Match struct {
	Id            int
	IssuerCatId   int
	ReceiverCatId int
	UserId        int
	Status        string
	Message       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
