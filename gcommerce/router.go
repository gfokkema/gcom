package gcommerce

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewRouter(endpoints Endpoints) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(contentType)
	r.Path("/articles/{id}").Handler(httptransport.NewServer(
		endpoints.GetArticle,
		decodeArticleReq,
		encodeResponse,
	)).Methods("GET")
	return r
}

func contentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
