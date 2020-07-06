package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yiplee/twirp-gateway-example/handler"
	"github.com/yiplee/twirp-gateway-example/handler/hc"
	"github.com/yiplee/twirp-gateway-example/store/book"
)

var cfg struct {
	port int
}

func main() {
	flag.IntVar(&cfg.port, "port", 8888, "server port")
	flag.Parse()

	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	books := book.New()
	svr := handler.New(books)

	r.Mount("/twirp", svr.HandleRpc())
	r.Mount("/api", svr.HandleApi())
	r.Mount("/hc", hc.Handler())

	addr := fmt.Sprintf(":%d", cfg.port)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
