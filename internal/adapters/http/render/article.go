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

type ArticleController struct {
    service port.IArticleService
}

func NewArticleController(service port.IArticleService) *ArticleController {
    return &ArticleController{
        service: service,
    }
}

func (c *ArticleController) HandleListArticles(w http.ResponseWriter, r *http.Request) {
    loadOrder := httpUtils.GetLoadOrder(r)    

    articleList := []*domain.Article{}
    var err error

    category := r.URL.Query().Get("category")
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
        fmt.Println("Error getting article list: ", err)
        w.WriteHeader(http.StatusInternalServerError)
    }

    fmt.Println("article list no controller de render: ", articleList)
    component := templates.ArticleList(articleList)

    var htmxHandler = HtmxHandler{
        Component: component,
    }

    htmxHandler.ServeHTTP(w, r)
}

func (c *ArticleController) GetArticleById(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")

    article, err := c.service.GetArticleById(r.Context(), id)
    if err != nil {
        fmt.Println("Error getting article: ", err)
        w.WriteHeader(http.StatusInternalServerError)
    }

    component := templates.Article(*article)

    var htmxHandler = HtmxHandler{
        Component: component,
    }

    htmxHandler.ServeHTTP(w, r )
}
