package gcommerce

type GcomSerializer interface {
	Decode(input []byte) (*Article, error)
	Encode(input *Article) ([]byte, error)
}
