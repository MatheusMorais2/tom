package port

import (
	"context"
	"tom/internal/core/domain"
)

type IKeywordRepository interface {
    CreateKeywords(ctx context.Context, keywords []*domain.Keyword) ([]*domain.Keyword, error)
    CreateArticleKeywords(ctx context.Context,postId string, keywords []*domain.Keyword) error
    GetKeywordsByArticleId(ctx context.Context, postId string) ([]*domain.Keyword, error)
}

type IKeywordService interface {
    CreateKeywords(ctx context.Context, keywords []*domain.Keyword) ([]*domain.Keyword, error)
    GetKeywordsByArticleId(ctx context.Context, postId string) ([]*domain.Keyword, error)
}
