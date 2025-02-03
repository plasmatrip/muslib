package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/plasmatrip/muslib/internal/model"
)

// AddSong добавляет новую песню
func (h *Handlers) AddSong(w http.ResponseWriter, r *http.Request) {
	var song model.Song

	// Разбираем тело запроса
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		h.Logger.Sugar.Infow("error in request handler", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверяем параметры
	if len(song.Song) == 0 || len(song.Group) == 0 {
		h.Logger.Sugar.Infow("error adding song", "error", errors.New("empty group name or song name"))
		http.Error(w, "empty group name or song name", http.StatusBadRequest)
		return
	}

	params := url.Values{}
	params.Add("group", song.Group)
	params.Add("song", song.Song)

	fullURL := fmt.Sprintf("%s/info?%s", h.Config.InfoService, params.Encode())

	// Отправляем запрос к внешнему сервису для получения раширенной информации о песне
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		h.Logger.Sugar.Infow("failed to create request", "error", err)
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		h.Logger.Sugar.Infow("failed to send request", "error", err)
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		h.Logger.Sugar.Infow("received non-200 response", "status: ", resp.StatusCode)
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}

	// Декодируем ответ
	if err := json.NewDecoder(resp.Body).Decode(&song); err != nil {
		h.Logger.Sugar.Infow("failed to decode response", "error", err)
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}

	// Добавляем песню в базу
	if err := h.Stor.AddSong(r.Context(), song); err != nil {
		h.Logger.Sugar.Infow("failed to add song", "error", err)
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}

	h.Logger.Sugar.Infow("song added successfully", "group", song.Group, "song", song.Song)

	w.WriteHeader(http.StatusOK)
}
