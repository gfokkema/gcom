package gcommerce

import (
	"errors"
	"time"
)

var (
	ErrArticleNotFound = errors.New("Article not found")
	ErrArticleInvalid  = errors.New("Article Invalid")
)

type gcomService struct {
	gcomRepo GcomRepository
}

func NewGcomService(gcomRepository GcomRepository) GcomService {
	return &gcomService{
		gcomRepository,
	}
}

func (g *gcomService) Find(id int64) (*Article, error) {
	return g.gcomRepo.Find(id)
}

func (g *gcomService) Store(article *Article) error {
	// maybe validate here
	article.CreatedAt = time.Now()
	return g.gcomRepo.Store(article)
}
