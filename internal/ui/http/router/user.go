package router

import (
	"github.com/go-chi/chi"
)

func (router *router) UserRouter() chi.Router {
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		// ユーザ情報取得
		r.Get("/user/{user_id}", router.handler.GetUser)
	})
	return r
}
