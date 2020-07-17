package gcommerce

type GcomRepository interface {
	Find(id int64) (*Article, error)
	Store(article *Article) error
}
