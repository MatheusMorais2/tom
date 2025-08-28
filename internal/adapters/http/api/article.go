package api

import (
	"encoding/json"
	"time"
	httpUtils "tom/internal/adapters/http/utils"
	"tom/internal/core/domain"
	"tom/internal/core/port"

	"net/http"

	"github.com/labstack/echo/v4"
)

type ArticleController struct {
    service port.IArticleService
}

func NewArticleController(service port.IArticleService) *ArticleController {
    return &ArticleController{
        service: service,
    }
}


func (c *ArticleController) GetLatestPosts(w http.ResponseWriter, r *http.Request) error {
    loadOrder := httpUtils.GetLoadOrder(r)
    articles, err := c.service.ListArticles(r.Context(), loadOrder.Skip, loadOrder.Limit)
    if err != nil {
        http.Error(w, "Could not get articles", 500)
        return err
    }

    jsonResponse, err := json.Marshal(articles)
    if err != nil {
        http.Error(w, "Could not parse json response", 500)
        return err
    }

    w.Header().Set("Content-type", "application-json")
    w.Write(jsonResponse)
    return nil
}

func (c *ArticleController) GetLatestAuditoriumPosts(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (c *ArticleController) GetLatestLoungePosts(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (c *ArticleController) GetLatestCinemaPosts(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (c *ArticleController) GetLatestBistroPosts(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (c *ArticleController) Post(w http.ResponseWriter, r *http.Request)  {
    ctx := r.Context()

    article := &domain.Article{}
    err := json.NewDecoder(r.Body).Decode(article) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    article, err = c.service.CreateArticle(ctx, article)
    if err != nil {
        // TODO: error middleware
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    w.Header().Set("Content-Type", "application/json")
    http.ResponseWriter.WriteHeader(w, http.StatusCreated)
    json.NewEncoder(w).Encode(article)

}

func (c *ArticleController) Patch(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    article := &domain.Article{}
    err := json.NewDecoder(r.Body).Decode(article) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    article, err = c.service.UpdateArticle(ctx, article)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    w.Header().Set("Content-Type", "application/json")
    http.ResponseWriter.WriteHeader(w, http.StatusCreated)
    json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) Delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type createPostRequest struct { 
    Id string `json:"id"`
    Title string `json:"title"`
    Summary string `json:"summary"`
    Content string  `json:"content"` //Markdown
    Category string `json:"category"`
    Keywords []string `json:"keywords"`
    CreatedAt time.Time `json:"created_at,omitmepty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (p *ArticleController) CreatePost (ctx echo.Context) error {
    var req *domain.Article
    if err := ctx.Bind(req); err != nil {
        return err
    }
    
    post, err := p.service.CreateArticle(ctx.Request().Context(), req)
    if err != nil {
        // TODO: error middleware
        return err
    }

    ctx.JSON(401, post)
    return nil
}
