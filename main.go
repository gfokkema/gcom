package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gfokkema/gcom/article"
	"github.com/gfokkema/gcom/catalog"
	"github.com/gfokkema/gcom/repository/inmem"
	"github.com/go-kit/kit/log"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
)

func NewArticleRepository(dsn string) article.Repository {
	// repo, err := sql.NewArticleRepository(dsn, time.Duration(5)*time.Second)
	// if err != nil {
	// 	panic(err)
	// }
	repo := inmem.NewArticleRepository()
	return repo
}

func main() {
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	articles := NewArticleRepository("gcom@/gcom?parseTime=true")

	svc := catalog.NewService(articles)

	mux := http.NewServeMux()
	mux.Handle("/catalog/v1/", catalog.MakeHandler(svc))
	http.Handle("/", mux)

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
