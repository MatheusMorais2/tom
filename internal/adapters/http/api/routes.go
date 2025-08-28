package api

import (
	"net/http"
	"tom/internal/core/port"
)

type Api struct {
    mux *http.ServeMux
}

func NewApi(mux *http.ServeMux) *Api {
    return &Api{
        mux: mux,
    }
}

func (api *Api) Router(articleController port.IArticleController) {
    api.mux.HandleFunc("POST /api/articles", articleController.Post)
    api.mux.HandleFunc("PATCH /api/articles", articleController.Patch)
}




