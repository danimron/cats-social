package repository

import (
	"cats_social/helper"
	"cats_social/model/domain"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type CatRepositoryImpl struct {
}

func NewCatRepository() CatRepository {
	return &CatRepositoryImpl{}
}

func (repository *CatRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cat domain.Cat) domain.Cat {
	sql := "INSERT INTO cats(user_id, name, race, sex, age_in_months, description, created_at, updated_at) values($1, $2, $3, $4, $5, $6, $7, $8)"
	cat.CreatedAt = time.Now()
	cat.UpdatedAt = time.Now()

	result, err := tx.ExecContext(ctx, sql, cat.UserId, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, cat.CreatedAt, cat.UpdatedAt)
	fmt.Println(err)

	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	cat.Id = int(id)
	return cat
}
