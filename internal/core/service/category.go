package service

import (
	"context"
	"fmt"
	"tom/internal/core/domain"
	"tom/internal/core/port"
)

type CategoryService struct {
	categoryRepo port.ICategoryRepository
}

func NewCategoryService(categoryRepo port.ICategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo}
}

func (cs *CategoryService) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
    if category == nil {
        return nil, nil
    }

    category, err := cs.categoryRepo.CreateCategory(ctx, category)
    if err != nil {
        fmt.Println("erro no repo: ", err)
        return nil, err
    }
    fmt.Printf("\n category: %+v", category)
	return category, nil
}

func (cs *CategoryService) GetCategoryId(ctx context.Context, name string) (*domain.Category, error) {
    return nil, nil
} 
