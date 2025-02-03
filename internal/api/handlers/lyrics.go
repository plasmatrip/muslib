package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/plasmatrip/muslib/internal/model"
)

// GetLyrics возвращает текст песни
func (h *Handlers) GetLyrics(w http.ResponseWriter, r *http.Request) {
	var song model.Song

	// Разбираем параметры
	query := r.URL.Query()

	if v := query.Get("group"); v != "" {
		song.Group = v
	}
	if v := query.Get("song"); v != "" {
		song.Song = v
	}

	verseNumStr := r.URL.Query().Get("verse")

	// Проверяем параметры
	verseNum, err := strconv.Atoi(verseNumStr)
	if err != nil || verseNum < 1 {
		h.Logger.Sugar.Infow("invalid verse number", "verse number", verseNumStr)
		http.Error(w, "invalid verse number", http.StatusBadRequest)
		return
	}

	// Получаем текст
	verse, err := h.Stor.GetLyrics(r.Context(), song, verseNum)
	if err != nil {
		h.Logger.Sugar.Infow("failed to fetch lyrics", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Logger.Sugar.Infow("got lyrics", "group", song.Group, "song", song.Song, "verse", verseNum)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(verse)
}
