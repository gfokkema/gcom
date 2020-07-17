package gcommerce

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetArticle endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetArticle: makeGetArticleEndpoint(s),
	}
}

func makeGetArticleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetArticleRequest)
		article, err := s.GetArticle(ctx, req.Id)
		return GetArticleResponse{Article: *article}, err
	}
}
