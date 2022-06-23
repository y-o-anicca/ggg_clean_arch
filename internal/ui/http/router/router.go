package router

import (
	"github.com/go-chi/chi"
	"github.com/y-o-anicca/ggg_clean_arch/internal/ui/http/handler"
	"github.com/y-o-anicca/ggg_clean_arch/internal/util/logger"
)

type (
	Router interface {
		Routing(r *chi.Mux)

		// user
		UserRouter() chi.Router
	}

	router struct {
		handler handler.Handler
		log     *logger.Logger
	}
)

func NewRouter(handler handler.Handler, log *logger.Logger) Router {
	return &router{
		handler,
		log,
	}
}

func (router *router) Routing(r *chi.Mux) {
	r.Get("/health", router.handler.Healthz)

	// user
	r.Mount("/user", router.UserRouter())
}
