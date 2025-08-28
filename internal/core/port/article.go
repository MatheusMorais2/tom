package port

import (
	"context"
	"net/http"
	"tom/internal/core/domain"
)

type IArticleRepository interface {
    CreateArticle(ctx context.Context, post *domain.Article) (*domain.Article, error)
    GetArticleById(ctx context.Context, id string) (*domain.Article, error)
    ListArticlesByCategory(ctx context.Context, category *domain.Category, skip, limit int) ([]*domain.Article, error)
    ListArticles(ctx context.Context, skip, limit int) ([]*domain.Article, error)
    UpdateArticle(ctx context.Context, post *domain.Article) (*domain.Article, error)
    DeleteArticle(ctx context.Context, id string) error 
}

type IArticleService interface {
    CreateArticle(ctx context.Context, post *domain.Article) (*domain.Article, error)
    GetArticleById(ctx context.Context, id string) (*domain.Article, error)
    ListArticles(ctx context.Context, skip, limit int) ([]*domain.Article, error)
    ListArticlesByCategory(ctx context.Context, category *domain.Category, skip, limit int) ([]*domain.Article, error)
    UpdateArticle(ctx context.Context, post *domain.Article) (*domain.Article, error)
    DeleteArticle(ctx context.Context, id string) error 
}

type IArticleController interface {
    Post(w http.ResponseWriter, r *http.Request)
    Patch(w http.ResponseWriter, r *http.Request)
}
