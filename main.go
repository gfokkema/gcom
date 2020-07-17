package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gfokkema/gcom/gcommerce"
	"github.com/gfokkema/gcom/repository/sql"
)

func NewRepo(dsn string) gcommerce.Repository {
	repo, err := sql.NewSqlRepository(dsn, time.Duration(5)*time.Second)
	if err != nil {
		panic(err)
	}
	return repo
}

func main() {
	repo := NewRepo("gcom@/gcom?parseTime=true")
	svc := gcommerce.NewService(repo)
	endpoints := gcommerce.MakeEndpoints(svc)
	router := gcommerce.NewRouter(endpoints)
	srv := &http.Server{Addr: ":8080", Handler: router}
	log.Fatal(srv.ListenAndServe())

	// article := &gcommerce.Article{
	// 	Title:     "goTest",
	// 	Desc:      "goDesc",
	// 	Price:     1.0,
	// 	CreatedAt: time.Now(),
	// }
	// err := svc.Store(article)
	// if err != nil {
	// 	panic(err)
	// }
}
