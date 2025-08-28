package render

import (
	"fmt"
	"net/http"
	httpUtils "tom/internal/adapters/http/utils"
	"tom/internal/core/domain"
	"tom/internal/core/port"
	"tom/internal/templates"
	//"tom/internal/core/service"
)

type ArticleListController struct {
    service port.IArticleService
}

func NewArticleListController(service port.IArticleService) *ArticleListController {
    return &ArticleListController{
        service: service,
    }
}

func (c *ArticleListController) HandleArticleList(w http.ResponseWriter, r *http.Request) {
    loadOrder := httpUtils.GetLoadOrder(r)    

    category := r.URL.Query().Get("category")

    articleList := []*domain.Article{}
    var err error

    switch category {
    case "auditorium":
        articleList, err = c.service.ListArticlesByCategory(r.Context(), 
            &domain.Category{Name: domain.Auditorium,}, 
            loadOrder.Skip, 
            loadOrder.Limit) 
    case "bistro":
        articleList, err = c.service.ListArticlesByCategory(r.Context(), 
            &domain.Category{Name: domain.Bistro,}, 
            loadOrder.Skip, 
            loadOrder.Limit) 
    
    case "cinema":
        articleList, err = c.service.ListArticlesByCategory(r.Context(), 
            &domain.Category{Name: domain.Cinema,}, 
            loadOrder.Skip, 
            loadOrder.Limit) 
    case "lounge":
        articleList, err = c.service.ListArticlesByCategory(r.Context(), 
            &domain.Category{Name: domain.Lounge,}, 
            loadOrder.Skip, 
            loadOrder.Limit) 
    default:
        articleList, err = c.service.ListArticles(r.Context(), 
            loadOrder.Skip, 
            loadOrder.Limit) 
    } 

    if err != nil {
        fmt.Errorf("Error getting article list: ", err)
    }

    component := templates.ArticleList(articleList)

    var htmxHandler = HtmxHandler{
        Component: component,
    }

    htmxHandler.ServeHTTP(w, r)
}
