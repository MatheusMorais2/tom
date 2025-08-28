package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"tom/internal/core/domain"
)

type KeywordRepository struct {
	db *sql.DB
}

func NewKeywordRepository(db *sql.DB) *KeywordRepository {
	return &KeywordRepository{
		db: db,
	}
}

func (kr *KeywordRepository) CreateKeywords(ctx context.Context, keywords []*domain.Keyword) ([]*domain.Keyword, error) {
    if len(keywords) == 0 {
        return keywords, nil
    }

    values := []interface{}{}  
    indexes := []string{}

    for i, keyword := range keywords {
        values = append(values, keyword.Name)
        indexes = append(indexes, fmt.Sprintf("($%d)", i+1))
    }
    query := fmt.Sprintf(`INSERT INTO keywords (name) VALUES %s ON CONFLICT (name) DO NOTHING;`, strings.Join(indexes, ", "))

	rows, err := kr.db.QueryContext(ctx, query, values...) 
	if err != nil {
		return nil, err
	}
    defer rows.Close()

    for i, index := range indexes {
        indexes[i] = strings.Trim(index, "()")
    }

    selectQuery := fmt.Sprintf(`SELECT id FROM keywords WHERE name IN (%s);`, strings.Join(indexes, ", "))
    
    selectRows, err := kr.db.QueryContext(ctx, selectQuery, values...)
	if err != nil {
		return nil, err
	}

    defer selectRows.Close()

    for i := 0; selectRows.Next(); i++ {
        if err := selectRows.Scan(&keywords[i].Id); err != nil {
            return nil, err
        } 
    }

	return keywords, nil
}

func (kr *KeywordRepository) CreateArticleKeywords(ctx context.Context, postId string, keywords []*domain.Keyword) error {
   if len(keywords) == 0 {
        return nil
   } 

   values := []interface{}{}
   indexes := []string{}
    values = append(values, postId)
   for i, keyword := range keywords {
        values = append(values, keyword.Id)
        indexes = append(indexes, fmt.Sprintf("($1, $%d)", i+2))
   }
   query := fmt.Sprintf(`INSERT INTO article_keywords (article_id, keyword_id) VALUES %s;`, strings.Join(indexes, ", "))

   _, err := kr.db.QueryContext(ctx, query, values...)
   if err != nil {
    return err
   }

   return nil
}

func (kr *KeywordRepository) GetKeywordsByArticleId(ctx context.Context, postId string) ([]*domain.Keyword, error) {
	query := `SELECT k.id, k.word 
        FROM keywords k
        INNER JOIN article_keywords ak ON ak.keyword_id = k.id
        WHERE ak.article_id = $1;`
	rows, err := kr.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keywords []*domain.Keyword
	for rows.Next() {
		var kw domain.Keyword
		if err := rows.Scan(&kw.Id, &kw.Name); err != nil {
			return nil, err
		}
		keywords = append(keywords, &kw)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return keywords, nil
}
