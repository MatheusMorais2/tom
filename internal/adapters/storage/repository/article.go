package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"tom/internal/core/domain"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (ar *ArticleRepository) CreateArticle(ctx context.Context, article *domain.Article) (*domain.Article, error) {
	fmt.Println("entrou no repo create article")
	query := `INSERT INTO articles (title, summary, content, category_id, author) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at;`
	err := ar.db.QueryRowContext(ctx, query, article.Title, article.Summary, article.Content, article.Category.Id, article.Author).Scan(&article.Id, &article.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("Error creating post: %+v", err)
	}

	fmt.Println("saiu no repo create article")
	return article, nil
}

func (ar *ArticleRepository) GetArticleById(ctx context.Context, id string) (*domain.Article, error) {
	article := &domain.Article{}

	query := `SELECT a.id, a.title, a.summary, a.content, c.id, c.name, a.created_at, a.updated_at
        FROM article a 
        LEFT JOIN categories c ON c.id = a.category_id
        WHERE a.id = $1`

	err := ar.db.QueryRowContext(ctx, query, id).Scan(
		&article.Id, &article.Title, &article.Summary, &article.Content, &article.Category.Id, &article.Category.Name, &article.CreatedAt, &article.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("Error getting post: %+v", err)
	}

	keywordsQuery := `SELECT k.id, k.name 
        FROM keywords k
        INNER JOIN article_keywords ak ON k.id = ak.keyword_id
        WHERE ak.article_id = $1
    `
	rows, err := ar.db.QueryContext(ctx, keywordsQuery, article.Id)
	if err != nil {
		return nil, fmt.Errorf("Error getting keywords: %+v", err)
	}

	defer rows.Close()

	for rows.Next() {
		kw := &domain.Keyword{}
		err := rows.Scan(&kw.Id, &kw.Name)
		if err != nil {
			return nil, fmt.Errorf("Error scanning keywords: %+v", err)
		}
		article.Keywords = append(article.Keywords, kw)
	}

	return article, nil
}

func (ar *ArticleRepository) ListArticles(ctx context.Context, skip, limit int) ([]*domain.Article, error) {
	query := `SELECT a.id, a.title, a.summary, c.id, c.name, a.created_at, a.updated_at
        FROM articles a
        LEFT JOIN categories c ON c.id = a.category_id
        ORDER BY created_at DESC
        LIMIT $1
        OFFSET $2`

	rows, err := ar.db.QueryContext(ctx, query, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("Error listing posts: %+v", err)
	}
	defer rows.Close()

	posts := []*domain.Article{}

	for rows.Next() {
		post := &domain.Article{}
		err := rows.Scan(&post.Id, &post.Title, &post.Summary, &post.Category.Id, &post.Category.Name, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Error scanning posts: %+v", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (ar *ArticleRepository) ListArticlesByCategory(ctx context.Context, category *domain.Category, skip, limit int) ([]*domain.Article, error) {
	query := `SELECT a.id, a.title, a.summary, a.author, a.created_at, a.updated_at
        FROM articles a 
        WHERE a.category_id = $3
        ORDER BY created_at DESC
        LIMIT $1
        OFFSET $2`

	rows, err := ar.db.QueryContext(ctx, query, limit, skip, category.Id)
	if err != nil {
		return nil, fmt.Errorf("Error listing posts: %+v", err)
	}
	defer rows.Close()

	articles := []*domain.Article{}

	for rows.Next() {
		article := &domain.Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Summary, &article.Author, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Error scanning posts: %+v", err)
		}
		article.Category = category
		articles = append(articles, article)
	}

	return articles, nil
}

func (ar *ArticleRepository) UpdateArticle(ctx context.Context, article *domain.Article) (*domain.Article, error) {
	query := `UPDATE articles SET`
	index := 1
	args := []interface{}{}

	fmt.Printf("\npost no repo: %+v", article)
	if article.Title != "" {
		args = append(args, article.Title)
		query, index = buildUpdateQuery("title", query, index)
	}

	if article.Summary != "" {
		args = append(args, article.Summary)
		query, index = buildUpdateQuery("summary", query, index)
	}

	if article.Content != "" {
		args = append(args, article.Content)
		query, index = buildUpdateQuery("content", query, index)
	}

	if article.Author != "" {
		args = append(args, article.Author)
		query, index = buildUpdateQuery("author", query, index)
	}

	if article.Category != nil {
		args = append(args, article.Category.Id)
		query, index = buildUpdateQuery("category_id", query, index)
	} else {
		article.Category = &domain.Category{}
	}

	args = append(args, time.Now())
	args = append(args, article.Id)

	query = fmt.Sprintf("%s, updated_at = $%d WHERE id = $%d RETURNING id, title, summary, content, author, category_id, created_at, updated_at", query, index, index+1)

	err := ar.db.QueryRowContext(ctx, query, args...).Scan(&article.Id, &article.Title, &article.Summary, &article.Content, &article.Author, &article.Category.Id, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("Error updating post: %+v", err)
	}

	return article, nil
}

func (ar *ArticleRepository) DeleteArticle(ctx context.Context, id string) error {
	query := `DELETE FROM articles WHERE id = $1`

	err := ar.db.QueryRowContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Error deleting post: %+v", err)
	}
	return nil
}

func buildUpdateQuery(column, query string, index int) (string, int) {
	if index != 1 {
		query = query + ","
	}
	query = fmt.Sprintf("%s %s = $%d", query, column, index)
	index++

	return query, index
}
