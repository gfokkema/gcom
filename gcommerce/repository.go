package gcommerce

type Repository interface {
	GetArticle(id int) (*Article, error)
	PostArticle(article *Article) error
}
