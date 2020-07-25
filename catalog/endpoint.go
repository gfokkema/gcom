package catalog

import (
	"context"
	"time"

	"github.com/gfokkema/gcom/article"
	"github.com/go-kit/kit/endpoint"
)

type (
	getArticleRequest struct {
		ID article.ArticleID `json:"id"`
	}
	getArticleResponse struct {
		Catalog article.Article `json:"article,omitempty"`
		Err     error           `json:"error,omitempty"`
	}
)

func (r getArticleResponse) error() error { return r.Err }

func makeGetArticleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getArticleRequest)
		catalog, err := s.GetArticle(req.ID)
		return getArticleResponse{Catalog: catalog, Err: err}, nil
	}
}

type (
	postArticleRequest struct {
		Title string  `json:"title"`
		Desc  string  `json:"desc"`
		Price float32 `json:"price"`
	}
	postArticleResponse struct {
		ID  article.ArticleID `json:"id,omitempty"`
		Err error             `json:"error,omitempty"`
	}
)

func (r postArticleResponse) error() error { return r.Err }

func makePostArticleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postArticleRequest)
		art := article.Article{
			Title:     req.Title,
			Desc:      req.Desc,
			Price:     req.Price,
			CreatedAt: time.Now(),
		}
		id, err := s.PostArticle(art)
		return postArticleResponse{ID: id, Err: err}, nil
	}
}
