package gcommerce

import "context"

type Service interface {
	GetArticle(ctx context.Context, id int) (*Article, error)
	PostArticle(ctx context.Context, article *Article) error
}
