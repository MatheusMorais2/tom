package port

import (
	"context"
	"tom/internal/core/domain"
)

type ICategoryRepository interface {
    CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
    GetCategoryId(ctx context.Context, name string) (*domain.Category, error) 
}

type ICategoryService interface {
    CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
    GetCategoryId(ctx context.Context, name string) (*domain.Category, error) 
}
