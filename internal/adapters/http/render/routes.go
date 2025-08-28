package render

import (
	"net/http"
	"time"
	"tom/internal/core/domain"
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

func (api *RenderApi) Router(articleListController *ArticleListController) {
	date := time.Now()
	var idk = domain.Article{
		Id:        "idman",
		Author:    "Matheus Morais",
		Content:   "Alo amor to the ligando d euma confusao ta um barulho uma confusao mas eu preciso tanto te falar depois das 3, to te esperando no meusmo lugar, to doido pra tee falar, viver amsi uma noite de aventura",
		Title:     "Como fazer essa merda funcionar",
		Summary:   "Essa porra eh uma porra pra resumir o caralho do caralho",
		CreatedAt: &date,
	}


	api.mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	api.mux.HandleFunc("GET /", HandleHomePage)
	api.mux.HandleFunc("GET /bistro", InitializeComponent(templates.Bistro()))
	api.mux.HandleFunc("GET /auditorium", InitializeComponent(templates.Auditorium()))
	api.mux.HandleFunc("GET /cinema", InitializeComponent(templates.Cinema()))
	api.mux.HandleFunc("GET /lounge", InitializeComponent(templates.Lounge()))
	api.mux.HandleFunc("GET /article", InitializeComponent(templates.Article(idk)))
	api.mux.HandleFunc("GET /article-list", articleListController.HandleArticleList)
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
