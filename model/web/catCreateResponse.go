package web

import (
	"time"
)

type CatCreateResponse struct {
	CatId     int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
