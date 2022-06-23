package handler

import (
	"net/http"
)

func (h *handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
