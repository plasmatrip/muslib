package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/plasmatrip/muslib/internal/model"
)

func (h *Handlers) UpdateSong(w http.ResponseWriter, r *http.Request) {
	var song model.Song

	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		h.Logger.Sugar.Infow("error in request handler", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(song.Song) == 0 || len(song.Group) == 0 {
		h.Logger.Sugar.Infow("error update song", "error", errors.New("empty group name or song name"))
		http.Error(w, "empty group name or song name", http.StatusBadRequest)
		return
	}

	if err := h.Stor.UpdateSong(r.Context(), song); err != nil {
		h.Logger.Sugar.Infow("failed to update song", "error", err)
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
