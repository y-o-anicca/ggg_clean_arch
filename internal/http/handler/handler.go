package handler

import (
	"net/http"

	"github.com/y-o-anicca/ggg_clean_arch/internal/usecase"
	"github.com/y-o-anicca/ggg_clean_arch/internal/util/logger"
)

type (
	Handler interface {
		Healthz(w http.ResponseWriter, r *http.Request)

		// user
		GetUser(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		useCase usecase.Usecase
		log     *logger.Logger
	}
)

func NewHandler(
	usecase usecase.Usecase,
	log *logger.Logger,
) Handler {
	return &handler{
		usecase,
		log,
	}
}
