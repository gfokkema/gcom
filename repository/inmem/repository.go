package inmem

import (
	"sync"

	"github.com/gfokkema/gcom/article"
	"github.com/pkg/errors"
)

type inMemRepository struct {
	mtx      sync.RWMutex
	articles map[article.ArticleID]article.Article
}

var (
	articleIdx = article.ArticleID(0)
)

func NewArticleRepository() article.Repository {
	return &inMemRepository{
		sync.RWMutex{},
		make(map[article.ArticleID]article.Article),
	}
}

func (repo *inMemRepository) Find(id article.ArticleID) (article.Article, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	v, ok := repo.articles[id]
	if !ok {
		return article.Article{}, errors.Wrap(article.ErrUnknown, "repository.Article.Find")
	}
	return v, nil
}

func (repo *inMemRepository) FindAll() ([]article.Article, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	values := make([]article.Article, len(repo.articles))
	for _, v := range repo.articles {
		values = append(values, v)
	}
	return values, nil
}

func (repo *inMemRepository) Store(art article.Article) (article.ArticleID, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	if _, ok := repo.articles[art.ID]; ok {
		return -1, errors.Wrap(article.ErrDuplicate, "repository.Article.Store")
	}

	articleIdx++
	art.ID = articleIdx
	repo.articles[articleIdx] = art

	return articleIdx, nil
}
