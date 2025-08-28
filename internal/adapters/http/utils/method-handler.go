package httpUtils

import "net/http"

type Controller interface {
	Get(w http.ResponseWriter, r *http.Request) 
	Post(w http.ResponseWriter, r *http.Request) 
	Delete(w http.ResponseWriter, r *http.Request) 
	Patch(w http.ResponseWriter, r *http.Request) 
}

// Eu quero passar uma request e quero receber um controler
func MethodHandler(c Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			c.Get(w, r)
		case http.MethodPost:
			c.Post(w, r)
		case http.MethodDelete:
			c.Delete(w, r)
		case http.MethodPatch:
			c.Patch(w, r)
		default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
