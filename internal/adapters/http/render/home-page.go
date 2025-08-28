package render

import (
	"net/http"
	"strconv"
	"tom/internal/templates"
	//"github.com/a-h/templ"
)

func HandleHomePage(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

    skip, err :=strconv.Atoi(query.Get("skip")) 
    limit, err := strconv.Atoi(query.Get("limit"))
    
    if err != nil {
        skip = 0
        limit = 10
    }

    loadOrder := templates.LoadOrder{
        Skip: skip,
        Limit: limit,
    }

    component := templates.HomePage(loadOrder)
    var htmxHandler = HtmxHandler{
        Component: component,
    }

    htmxHandler.ServeHTTP(w, r)
}
