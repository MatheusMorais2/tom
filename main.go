package main

import (
	"fmt"
	"net/http"
	"github.com/a-h/templ"
    "tom/pkg/templates"
)

func main() {
    component := templates.Index()
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) 
    http.Handle("/", templ.Handler(component))
    http.Handle("/home-page", templ.Handler(templates.Content()))
    http.Handle("/bistro", templ.Handler(templates.Bistro()))
    http.Handle("/auditorium", templ.Handler(templates.Auditorium()))
    http.Handle("/cinema", templ.Handler(templates.Cinema()))
    http.Handle("/lounge", templ.Handler(templates.Lounge()))

    httpPort := ":3000"
    fmt.Println("Listening on port ", httpPort)
    http.ListenAndServe(httpPort, nil)
}
