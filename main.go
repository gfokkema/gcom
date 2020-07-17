package main

import (
	"fmt"
	"time"

	"github.com/gfokkema/gcom/gcommerce"
	"github.com/gfokkema/gcom/repository/sql"
)

func main() {
	repo, err := sql.NewSqlRepository("gcom@/gcom?parseTime=true", time.Duration(5)*time.Second)
	if err != nil {
		panic(err)
	}
	svc := gcommerce.NewGcomService(repo)

	article := &gcommerce.Article{
		Title:     "goTest",
		Desc:      "goDesc",
		Price:     1.0,
		CreatedAt: time.Now(),
	}
	err = svc.Store(article)
	if err != nil {
		panic(err)
	}
	article, err = svc.Find(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", article)
}
