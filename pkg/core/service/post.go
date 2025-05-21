package service

import (
	"context"
	"tom/pkg/core/domain"
	"tom/pkg/core/port"
)

type PostService struct {
    repo port.PostRepository
}

func NewPostService(repo *port.PostRepository) *PostService {
    return &PostService{
        repo: *repo,
    }
}

func (ps *PostService) CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error) {
    p, err := ps.repo.CreatePost(ctx, post)
    if err != nil {
        return nil, err
    }
    return p, nil
}

func (ps *PostService) GetPostById(ctx context.Context, id string) (*domain.Post, error) {
    p, err := ps.repo.GetPostById(ctx, id)
    if err != nil {
        return nil, err
    }
    return p, nil
}

func (ps *PostService) ListPosts(ctx context.Context, skip, limit int) ([]*domain.Post, error) {
    posts, err := ps.repo.ListPosts(ctx, skip, limit)
    if err != nil {
        return nil, err
    }
    return posts, nil
}

func (ps *PostService) ListPostsByCategory(ctx context.Context, category domain.Category, skip, limit int) ([]*domain.Post, error) {
    posts, err := ps.repo.ListPostsByCategory(ctx, category, skip, limit)
    if err != nil {
        return nil, err
    }
    return posts, nil
}

func (ps *PostService) UpdatePost(ctx context.Context, post *domain.Post) (*domain.Post, error) {
    p, err := ps.repo.UpdatePost(ctx, post)
    if err != nil {
        return nil, err
    }
    return p, nil
}

func (ps *PostService) DeletePost(ctx context.Context, id string) error {
    err := ps.repo.DeletePost(ctx, id)
    if err != nil {
        return err
    }
    return nil
}
