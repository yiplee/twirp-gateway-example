package hc

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yiplee/twirp-gateway-example/handler/render"
)

func Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.NoCache)
	r.Handle("/", handle())
	return r
}

func handle() http.HandlerFunc {
	start := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(start).Truncate(time.Millisecond)

		render.JSON(w, render.H{
			"uptime": uptime.String(),
		})
	}
}
