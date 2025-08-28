package repository

import (
	"context"
	"database/sql"
	"fmt"
	"tom/internal/core/domain"
)

type CategoryRepository struct {
    db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
    return &CategoryRepository{
        db: db,
    }
}

func (cr *CategoryRepository) GetCategoryId(ctx context.Context, name string) (*domain.Category, error) {
    category := &domain.Category{}
    query := `SELECT id, name FROM categories WHERE name = $1;`
    err := cr.db.QueryRowContext(ctx, query, name).Scan(&category.Id, &category.Name)
    if err != nil {
        return nil, err
    }

    return category, nil
}

func (cr *CategoryRepository) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
    fmt.Printf("\n name: %s" , category.Name)
    query := `INSERT INTO categories (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=$1 returning id`
    err := cr.db.QueryRowContext(ctx, query, category.Name).Scan(&category.Id)
    if err != nil {
        return nil, err
    }
	return category, nil
}
