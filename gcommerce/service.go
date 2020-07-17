package gcommerce

type GcomService interface {
	Find(id int64) (*Article, error)
	Store(article *Article) error
}
