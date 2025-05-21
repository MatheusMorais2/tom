package api

import (
	"time"
	"tom/pkg/core/domain"
	"tom/pkg/core/port"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
    service port.PostService
}

func NewPostHandler(service *port.PostService) *PostHandler {
    return &PostHandler{
        service: *service,
    }
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

func (p *PostHandler) CreatePost (ctx echo.Context) error {
    var req *domain.Post
    if err := ctx.Bind(req); err != nil {
        return err
    }
    
    post, err := p.service.CreatePost(ctx.Request().Context(), req)
    if err != nil {
        // TODO: error middleware
        return err
    }

    ctx.JSON(401, post)
    return nil
}
