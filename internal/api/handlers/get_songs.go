package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/plasmatrip/muslib/internal/model"
)

func (h *Handlers) GetSongs(w http.ResponseWriter, r *http.Request) {
	// Разбираем параметры запроса
	filter, err := parseQueryParams(r)
	if err != nil {
		h.Logger.Sugar.Infow("failed to parse query params", "error", err)
		http.Error(w, "invalid query parameters", http.StatusBadRequest)
		return
	}

	// Достаем данные из базы
	songs, err := h.Stor.GetSongs(r.Context(), filter)
	if err != nil {
		h.Logger.Sugar.Infow("failed to fetch songs", "error", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if len(songs) == 0 {
		h.Logger.Sugar.Debugw("no songs found. filter:", "group", filter.Group,
			"song", filter.Song, "text", filter.Text, "link", filter.Link, "release_from", filter.ReleaseFrom, "release_to", filter.ReleaseTo)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

func parseQueryParams(r *http.Request) (*model.Filter, error) {
	filter := &model.Filter{
		Limit:  10,
		Offset: 0,
	}

	query := r.URL.Query()

	if v := query.Get("group"); v != "" {
		filter.Group = &v
	}
	if v := query.Get("song"); v != "" {
		filter.Song = &v
	}
	if v := query.Get("text"); v != "" {
		filter.Text = &v
	}
	if v := query.Get("link"); v != "" {
		filter.Link = &v
	}

	if v := query.Get("release_from"); v != "" {
		t, err := time.Parse("02-01-2006", v)
		if err != nil {
			return nil, err
		}
		filter.ReleaseFrom = &t
	}
	if v := query.Get("release_to"); v != "" {
		t, err := time.Parse("02-01-2006", v)
		if err != nil {
			return nil, err
		}
		filter.ReleaseTo = &t
	}

	if v := query.Get("limit"); v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil || limit <= 0 {
			return nil, err
		}
		filter.Limit = limit
	}
	if v := query.Get("offset"); v != "" {
		offset, err := strconv.Atoi(v)
		if err != nil || offset < 0 {
			return nil, err
		}
		filter.Offset = offset
	}

	return filter, nil
}
