package json

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Article struct{}

func (r *Article) Decode(input []byte) (*gcommerce.Article, error) {
	article := &gcommerce.Article{}
	if err := json.Unmarshal(input, article); err != nil {
		return nil, errors.Wrap(err, "serializer.Article.Decode")
	}
	return article, nil
}

func (r *Article) Encode(input *Article) ([]byte, error) {
	msg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Article.Encode")
	}
	return msg, nil
}
