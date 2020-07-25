package article

import (
	"errors"
	"time"
)

var (
	ErrUnknown   = errors.New("Unknown article")
	ErrDuplicate = errors.New("Duplicate article")
)

type ArticleID int32

type Article struct {
	ID        ArticleID `db:"id"`
	Title     string    `db:"title"`
	Desc      string    `db:"description"`
	Price     float32   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}

// Repository provides access a cargo store.
type Repository interface {
	Find(id ArticleID) (Article, error)
	FindAll() ([]Article, error)
	Store(article Article) (ArticleID, error)
}
