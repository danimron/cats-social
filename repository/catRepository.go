package repository

import (
	"cats_social/model/domain"
	"context"
	"database/sql"
)

type CatRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cat domain.Cat) domain.Cat
	// Delete(ctx context.Context, tx *sql.Tx, cat domain.Cat)
	// Update(ctx context.Context, tx *sql.Tx, cat domain.Cat) domain.Cat
	// FindAll(ctx context.Context, db *sql.DB) []domain.Cat
}
