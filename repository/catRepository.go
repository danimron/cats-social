package repository

import (
	"cats_social/model/domain"
	"context"
	"database/sql"
)

type CatRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cat domain.Cat) domain.Cat
	Delete(ctx context.Context, tx *sql.Tx, catId int)
	Update(ctx context.Context, tx *sql.Tx, cat domain.Cat)
	FindById(ctx context.Context, db *sql.DB, catId int) (domain.Cat, error)
	FindAll(ctx context.Context, db *sql.DB, sql string) []domain.Cat
}
