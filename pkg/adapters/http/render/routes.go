package routes

import (
	"net/http"
	"tom/pkg/templates"

	"github.com/a-h/templ"
)

func Routes() {
    http.Handle("/bistro", templ.Handler(templates.Bistro()))
    http.Handle("/auditorium", templ.Handler(templates.Auditorium()))
    http.Handle("/cinema", templ.Handler(templates.Cinema()))
    http.Handle("/lounge", templ.Handler(templates.Lounge()))
}
