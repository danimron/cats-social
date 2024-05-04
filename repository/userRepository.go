package repository

import (
	"cats_social/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cat domain.User) (domain.User, error)
	// Delete(ctx context.Context, tx *sql.Tx, cat domain.Cat)
	// Update(ctx context.Context, tx *sql.Tx, cat domain.Cat) domain.Cat
	// FindAll(ctx context.Context, db *sql.DB) []domain.Cat
}
