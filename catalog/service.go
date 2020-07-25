package catalog

import (
	"github.com/gfokkema/gcom/article"
)

type Service interface {
	GetArticle(article.ArticleID) (article.Article, error)
	GetArticleAll() ([]article.Article, error)
	PostArticle(article.Article) (article.ArticleID, error)
}

type service struct {
	articles article.Repository
}

// NewService creates a booking service with necessary dependencies.
func NewService(articles article.Repository) Service {
	return &service{
		articles: articles,
	}
}

func (s *service) GetArticle(ID article.ArticleID) (article.Article, error) {
	return s.articles.Find(ID)
}

func (s *service) GetArticleAll() ([]article.Article, error) {
	return s.articles.FindAll()
}

func (s *service) PostArticle(a article.Article) (article.ArticleID, error) {
	return s.articles.Store(a)
}
