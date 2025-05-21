package repository

import (
	"context"
	"database/sql"
	"tom/pkg/core/domain"
)

type PostRepository struct {
    db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
    return &PostRepository{
        db: db,
    }
}

func (pr *PostRepository) CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error) {
    return nil, nil
}

func (pr *PostRepository) GetPostById(ctx context.Context, id string) (*domain.Post, error) {
    return nil, nil
}

func (pr *PostRepository) ListPosts(ctx context.Context, skip, limit int) (*domain.Post, error) {
    return nil, nil
}

func (pr *PostRepository) ListPostsByCategory(ctx context.Context, category domain.Category, skip, limit int) (*domain.Post, error) {
    return nil, nil
}

func (pr *PostRepository) UpdatePost(ctx context.Context, post *domain.Post) (*domain.Post, error) {
    return nil, nil
}

func (pr *PostRepository) DeletePost(ctx context.Context, id string) error {
    return nil
}
