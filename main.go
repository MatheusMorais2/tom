package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tom/internal/adapters/http/api"
	"tom/internal/adapters/http/render"
	"tom/internal/adapters/storage"
	"tom/internal/adapters/storage/repository"
	"tom/internal/core/service"

	/*"context"
		"tom/internal/adapters/storage/repository"
		"tom/internal/core/service"
	    "time"
		"tom/internal/core/domain"
	*/
	"tom/internal/templates"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := storage.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	categoryRepo := repository.NewCategoryRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	keywordRepo := repository.NewKeywordRepository(db)

    _ = service.NewCategoryService(categoryRepo)
    articleService := service.NewArticleService(articleRepo, keywordRepo, categoryRepo)
     _ = service.NewKeywordService(keywordRepo)
    
    articleController := api.NewArticleController(articleService)
    renderArticleController := render.NewArticleController(articleService)

    mux := http.NewServeMux()

    renderApi := render.NewRenderApi(mux)
    renderApi.Router(renderArticleController)

    API := api.NewApi(mux)
    API.Router(articleController)

    httpPort := ":" + os.Getenv("PORT")
	fmt.Println("Listening on port ", httpPort)
	http.ListenAndServe(httpPort, mux)
}

type HtmxHandler struct {
	Component templ.Component
}

func (hx HtmxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") != "" {
		templ.Handler(hx.Component).ServeHTTP(w, r)
	} else {
		templ.Handler(templates.Index(hx.Component)).ServeHTTP(w, r)
	}
}

func InitializeComponent(component templ.Component) http.HandlerFunc {
	var htmxHandler = HtmxHandler{
		Component: component,
	}

	return htmxHandler.ServeHTTP
}
