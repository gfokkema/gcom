package sql

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gfokkema/gcom/article"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type sqlRepository struct {
	client  *sqlx.DB
	dsn     string
	timeout time.Duration
}

func newSQLClient(dsn string, timeout time.Duration) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = client.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewArticleRepository(dsn string, timeout time.Duration) (article.Repository, error) {
	client, err := newSQLClient(dsn, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewSqlRepo")
	}
	return &sqlRepository{
		client:  client,
		dsn:     dsn,
		timeout: timeout,
	}, nil
}

func (repo *sqlRepository) Find(id article.ArticleID) (article.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	article := article.Article{}
	err := repo.client.GetContext(ctx, &article, "SELECT * FROM articles WHERE id = ?", id)
	if err != nil {
		return article, errors.Wrap(err, "repository.Article.Find")
	}
	return article, nil
}

func (repo *sqlRepository) FindAll() ([]article.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	articles := []article.Article{}
	err := repo.client.SelectContext(ctx, articles, "SELECT * FROM articles")
	if err != nil {
		return nil, errors.Wrap(err, "repository.Article.FindAll")
	}
	return articles, nil
}

func (repo *sqlRepository) Store(art article.Article) (article.ArticleID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	v, err := repo.client.NamedExecContext(
		ctx,
		`INSERT INTO articles
			(title, description, price, created_at)
			VALUES
			(:title, :description, :price, :created_at)
		`,
		art,
	)
	if err != nil {
		return -1, errors.Wrap(err, "repository.Article.Store")
	}

	k, err := v.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "repository.Article.Store")
	}

	return article.ArticleID(k), nil
}
