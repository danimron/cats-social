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

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user := domain.User{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
	}
	user = service.UserRepository.Save(ctx, tx, user)
	return helper.ToCategoryResponseUser(user)
}
