package gcommerce

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	GetArticleRequest struct {
		Id int `json:"id"`
	}
	GetArticleResponse struct {
		Article Article `json:"article"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeArticleReq(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	return GetArticleRequest{
		Id: id,
	}, nil
}
