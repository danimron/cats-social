package service

import (
	"cats_social/model/web"
	"context"
)

type CatService interface {
	Create(ctx context.Context, request web.CatCreateRequest) web.CatCreateResponse
	// Update(ctx context.Context, request web.CatUpdateRequest) web.CatUpdateResponse
	// Delete(ctx context.Context, CatId int)
	// FindAll(ctx context.Context) []web.CatGetResponse
}
