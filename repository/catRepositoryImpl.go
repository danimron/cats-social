package repository

import (
	"cats_social/helper"
	"cats_social/model/domain"
	"context"
	"database/sql"
	"time"
)

type CatRepositoryImpl struct {
}

func NewCatRepository() CatRepository {
	return &CatRepositoryImpl{}
}

func (repository *CatRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cat domain.Cat) domain.Cat {
	sql := "INSERT INTO cats(user_id, name, race, sex, age_in_months, description, created_at, updated_at) values($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, cat.UserId, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, time.Now(), time.Now()).Scan(&insertedId)
	helper.PanicIfError(err)
	cat.Id = insertedId
	return cat
}

func (repository *CatRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, catId int) {
	sql := "UPDATE cats SET deleted_at = $1, updated_at = $2 WHERE id = $3"
	_, err := tx.ExecContext(ctx, sql, time.Now(), time.Now(), catId)
	helper.PanicIfError(err)
}

func (repository *CatRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, cat domain.Cat) {
	sql := "UPDATE cats SET name = $1, race = $2, sex = $3, age_in_months = $4, description = $5, updated_at = $6 WHERE id = $7"
	_, err := tx.ExecContext(ctx, sql, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, time.Now(), cat.Id)
	helper.PanicIfError(err)
}

func (repository *CatRepositoryImpl) FindById(ctx context.Context, db *sql.DB, catId int) (domain.Cat, error) {
	sql := "SELECT * FROM cats WHERE id = $1"
	rows, err := db.QueryContext(ctx, sql, catId)
	helper.PanicIfError(err)
	cat := domain.Cat{}
	defer rows.Close()
	if !rows.Next() {
		err := rows.Scan(&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.CreatedAt)
		helper.PanicIfError(err)
		return cat, nil
	} else {
		// return cat, errors.New("book is not found")
		return cat, nil
	}
}

func (repository *CatRepositoryImpl) FindAll(ctx context.Context, db *sql.DB, sql string) []domain.Cat {
	rows, err := db.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	var cats []domain.Cat
	for rows.Next() {
		cat := domain.Cat{}
		err := rows.Scan(&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.CreatedAt)
		helper.PanicIfError(err)
		cats = append(cats, cat)
	}
	return cats
}
