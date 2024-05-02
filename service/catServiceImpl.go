package service

import (
	// "cats_social/exception"
	"cats_social/helper"
	"cats_social/model/domain"
	"cats_social/model/web"
	"cats_social/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CatServiceImpl struct {
	CatRepository repository.CatRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewCatService(catRepository repository.CatRepository, DB *sql.DB, validate *validator.Validate) CatService {
	return &CatServiceImpl{
		CatRepository: catRepository,
		DB:            DB,
		Validate:      validate,
	}
}

func (service *CatServiceImpl) Create(ctx context.Context, request web.CatCreateRequest) web.CatCreateResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	cat := domain.Cat{
		UserId:      request.UserId,
		Name:        request.Name,
		Race:        request.Race,
		Sex:         request.Sex,
		AgeInMonth:  request.AgeInMonth,
		Description: request.Description,
	}
	cat = service.CatRepository.Save(ctx, tx, cat)
	return helper.ToCategoryResponseCat(cat)
}
