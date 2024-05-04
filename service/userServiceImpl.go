package service

import (
	"cats_social/helper"
	"cats_social/model/domain"
	"cats_social/model/web"
	"cats_social/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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

// func (service *UserServiceImpl) GenerateToken(ctx context.Context, student domain.User) string {
// 	expTime := time.Now().Add(time.Minute * 2)
// 	claims := &config.JWTClaim{
// 		Name:      student.Name,
// 		StudentId: student.StudentId,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expTime),
// 		},
// 	}
// 	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	key := []byte(os.Getenv("JWT_KEY"))
// 	token, err := generateToken.SignedString(key)
// 	helper.PanicIfError(err)
// 	return token
// 	// set cookie
// }

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	request.Password = string(bytes)

	user := domain.User{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
	}
	user, err = service.UserRepository.Save(ctx, tx, user)
	if err != nil {
		return web.UserResponse{}, err
	}
	return helper.ToCategoryResponseUser(user), nil
}
