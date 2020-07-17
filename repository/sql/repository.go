package sql

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gfokkema/gcom/gcommerce"
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

func NewSqlRepository(dsn string, timeout time.Duration) (gcommerce.Repository, error) {
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

func (r *sqlRepository) GetArticle(id int) (*gcommerce.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	article := &gcommerce.Article{}
	err := r.client.GetContext(ctx, article, "SELECT * FROM articles WHERE id = ?", id)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Article.Find")
	}
	return article, nil
}

func (r *sqlRepository) PostArticle(article *gcommerce.Article) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	_, err := r.client.NamedExecContext(
		ctx,
		`INSERT INTO articles
			(title, description, price, created_at)
			VALUES
			(:title, :description, :price, :created_at)
		`,
		article,
	)
	return err
}
