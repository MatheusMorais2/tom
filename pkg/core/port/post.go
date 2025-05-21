package port

import (
	"context"
	"tom/pkg/core/domain"
)

type PostRepository interface {
    CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error)
    GetPostById(ctx context.Context, id string) (*domain.Post, error)
    ListPostsByCategory(ctx context.Context, category domain.Category, skip, limit int) ([]*domain.Post, error)
    ListPosts(ctx context.Context, skip, limit int) ([]*domain.Post, error)
    UpdatePost(ctx context.Context, post *domain.Post) (*domain.Post, error)
    DeletePost(ctx context.Context, id string) error 
}

type PostService interface {
    CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error)
    GetPostById(ctx context.Context, id string) (*domain.Post, error)
    ListPostsByCategory(ctx context.Context, category domain.Category, skip, limit int) ([]*domain.Post, error)
    ListPosts(ctx context.Context, skip, limit int) ([]*domain.Post, error)
    UpdatePost(ctx context.Context, post *domain.Post) (*domain.Post, error)
    DeletePost(ctx context.Context, id string) error 
}
