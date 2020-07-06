package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yiplee/twirp-gateway-example/core"
	"github.com/yiplee/twirp-gateway-example/handler/render"
	"github.com/yiplee/twirp-gateway-example/handler/reversetwirp"
	"github.com/yiplee/twirp-gateway-example/handler/rpc/book"
)

func New(
	books core.BookStore,
) Server {
	return Server{
		Books: books,
	}
}

type Server struct {
	Books core.BookStore
}

func (s Server) HandleRpc() http.Handler {
	r := chi.NewRouter()
	r.Use(resetRoutePath)

	// book service
	{
		svc := book.New(s.Books)
		r.Mount(svc.PathPrefix(), svc)
	}

	return r
}

func (s Server) HandleApi() http.Handler {
	r := chi.NewRouter()
	r.Use(render.WrapResponse(true))

	r.Route("/books", func(r chi.Router) {
		svc := book.New(s.Books)
		rt := reversetwirp.NewSingleTwirpServerProxy(svc)

		r.Post("/", rt.Handle("Create", nil))
		r.Get("/{id}", rt.Handle("Find", nil))
	})

	return r
}

func resetRoutePath(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if c := chi.RouteContext(ctx); c != nil {
			c.RoutePath = r.URL.Path
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
