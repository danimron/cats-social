package service

import (
	"cats_social/model/web"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error)
	// Update(ctx context.Context, request web.CatUpdateRequest) web.CatUpdateResponse
	// Delete(ctx context.Context, CatId int)
	// FindAll(ctx context.Context) []web.CatGetResponse
}
