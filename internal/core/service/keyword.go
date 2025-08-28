package service

import (
	"context"
	"fmt"
	"strings"
	"tom/internal/core/domain"
	"tom/internal/core/port"
)

type KeywordService struct {
    repo port.IKeywordRepository
}

func NewKeywordService(repo port.IKeywordRepository) *KeywordService {
    return &KeywordService{
        repo: repo,
    }
}

func (ks *KeywordService) CreateKeywords(ctx context.Context, keywords []*domain.Keyword) ([]*domain.Keyword, error) {
    fmt.Println("chegou no crete keyword service")
    for _, keyword := range keywords {
        keyword.Name = normalizeKeyword(keyword.Name)
    }
    fmt.Print(keywords)   
    keywords, err := ks.repo.CreateKeywords(ctx, keywords)
    if err != nil {
        return nil, err
    }
    return keywords, nil
}

//TODO: remove if not used
func (ks *KeywordService) GetKeywordsByPostId(ctx context.Context, postId string) ([]*domain.Keyword, error) {
    return nil, nil
}

func normalizeKeyword(keyword string) string {
    keyword = strings.ToLower(keyword)
    keyword = strings.TrimSpace(keyword)
    return keyword
}

func createPostKeywords() {
    
}
