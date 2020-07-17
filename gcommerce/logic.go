package gcommerce

import (
	"context"
	"errors"
	"time"
)

var (
	ErrArticleNotFound = errors.New("Article not found")
	ErrArticleInvalid  = errors.New("Article Invalid")
)

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) GetArticle(ctx context.Context, id int) (*Article, error) {
	return s.repo.GetArticle(id)
}

func (s *service) PostArticle(ctx context.Context, article *Article) error {
	// maybe validate here
	article.CreatedAt = time.Now()
	return s.repo.PostArticle(article)
}
