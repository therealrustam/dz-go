package verify

import "net/http"

type Handler struct{}

func NewHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("GET /verify/{hash}", handler.VerifyEmail())
	router.HandleFunc("POST /send", handler.PostEmail())
}

func (handler *Handler) VerifyEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}

}

func (handler *Handler) PostEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
