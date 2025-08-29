package render

import (
	"net/http"
	"tom/internal/templates"

	"github.com/a-h/templ"
)

type RenderApi struct {
    mux *http.ServeMux
}

func NewRenderApi(mux *http.ServeMux) *RenderApi {
    return &RenderApi{
        //eu poderia colocar os controllers que eu tenho que injetar aqui, nao?
        mux: mux,
    }
}

func (api *RenderApi) Router(articleController *ArticleController) {
	api.mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	api.mux.HandleFunc("GET /", HandleHomePage)
	api.mux.HandleFunc("GET /bistro", InitializeComponent(templates.Bistro()))
	api.mux.HandleFunc("GET /auditorium", InitializeComponent(templates.Auditorium()))
	api.mux.HandleFunc("GET /cinema", InitializeComponent(templates.Cinema()))
	api.mux.HandleFunc("GET /lounge", InitializeComponent(templates.Lounge()))
	api.mux.HandleFunc("GET /article", articleController.GetArticleById)
	api.mux.HandleFunc("GET /article-list", articleController.HandleListArticles)
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
