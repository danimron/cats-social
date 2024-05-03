package repository

import (
	"cats_social/helper"
	"cats_social/model/domain"
	"context"
	"database/sql"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := "INSERT INTO users(name,email,password, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&insertedId)
	helper.PanicIfError(err)
	user.Id = insertedId
	return user
}
