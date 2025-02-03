package handlers

import (
	"net/http"
)

func (h *Handlers) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := h.Stor.Ping(r.Context())
	if err != nil {
		h.Logger.Sugar.Infow(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
