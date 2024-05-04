package repository

import (
	"cats_social/helper"
	"cats_social/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	var exists bool
	errValidation := tx.QueryRow("SELECT exists(SELECT 1 FROM users WHERE email=$1)", user.Email).Scan(&exists)
	if errValidation != nil {
		fmt.Print("Email not found")
	}
	if exists {
		fmt.Printf("Email found %s", user.Email)
		return user, errors.New("Email already exists")
	}
	sql := "INSERT INTO users(name,email,password, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&insertedId)
	helper.PanicIfError(err)
	user.Id = insertedId
	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	var user domain.User
	sql := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email=$1"
	err := tx.QueryRowContext(ctx, sql, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, errors.New("User not found")
	}
	return user, nil
}
