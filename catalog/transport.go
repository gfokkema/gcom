package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gfokkema/gcom/article"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	getArticleHandler := kithttp.NewServer(
		makeGetArticleEndpoint(s),
		decodeGetArticleRequest,
		encodeResponse,
		opts...,
	)
	postArticleHandler := kithttp.NewServer(
		makePostArticleEndpoint(s),
		decodePostArticleRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/catalog/v1/articles/{id}", getArticleHandler).Methods("GET")
	r.Handle("/catalog/v1/articles/", postArticleHandler).Methods("POST")
	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, errBadRoute
	}
	return getArticleRequest{ID: article.ArticleID(id)}, nil
}

func decodePostArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := postArticleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
