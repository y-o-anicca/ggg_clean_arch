package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type GetUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil || id < 1 {
		return
	}
	user, err := h.useCase.GetUser(ctx, id)
	if err != nil {
		return
	}

	render.JSON(w, r, GetUserResponse{
		ID:   user.ID,
		Name: user.Name,
	})
}
