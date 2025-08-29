package service

import (
	"context"
	"fmt"
	"tom/internal/core/domain"
	"tom/internal/core/port"
)

type ArticleService struct {
	articleRepo port.IArticleRepository
	keywordRepo port.IKeywordRepository
	categoryRepo port.ICategoryRepository
}

func NewArticleService(postRepo port.IArticleRepository, keywordRepo port.IKeywordRepository, categoryRepo port.ICategoryRepository) *ArticleService {
	return &ArticleService{
		articleRepo: postRepo,
		keywordRepo: keywordRepo,
        categoryRepo: categoryRepo,
	}
}

func (as *ArticleService) CreateArticle(ctx context.Context, article *domain.Article) (*domain.Article, error) {
	fmt.Printf("\narticle:  %+v", article)
    fmt.Printf("\nkeyword[0]: %+v", article.Keywords[0])

	keywordService := NewKeywordService(as.keywordRepo)
	keywords, err := keywordService.CreateKeywords(ctx, article.Keywords)
	if err != nil {
		return nil, err
	}
	article.Keywords = keywords

    CategoryService := NewCategoryService(as.categoryRepo)
    category, err := CategoryService.CreateCategory(ctx, article.Category)
	if err != nil {
		return nil, err
	}
    article.Category = category

	p, err := as.articleRepo.CreateArticle(ctx, article)
	if err != nil {
		return nil, err
	}


	if err = keywordService.repo.CreateArticleKeywords(ctx, p.Id, keywords); err != nil {
		fmt.Printf("\n erro criando post keywords: %+v: ", err)
		return nil, err
	}

    fmt.Println("deu bom: ", p)
	return p, nil
}

func (as *ArticleService) GetArticleById(ctx context.Context, id string) (*domain.Article, error) {
	p, err := as.articleRepo.GetArticleById(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (as *ArticleService) ListArticles(ctx context.Context, skip, limit int) ([]*domain.Article, error) {
	fmt.Printf("\nchegou no service list post com: %v e %v", skip, limit)
	articles, err := as.articleRepo.ListArticles(ctx, skip, limit)
	if err != nil {
		return nil, err
	}
    fmt.Println("saiu do service em")
	return articles, nil
}

func (as *ArticleService) ListArticlesByCategory(ctx context.Context, category *domain.Category, skip, limit int) ([]*domain.Article, error) {
	fmt.Printf("\nchegou no service list post com: %v e %v", skip, limit)
	posts, err := as.articleRepo.ListArticlesByCategory(ctx, category, skip, limit)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (as *ArticleService) UpdateArticle(ctx context.Context, post *domain.Article) (*domain.Article, error) {
	p, err := as.articleRepo.UpdateArticle(ctx, post)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (as *ArticleService) DeleteArticle(ctx context.Context, id string) error {
	err := as.articleRepo.DeleteArticle(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
