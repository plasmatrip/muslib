package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/plasmatrip/muslib/internal/model"
)

func (h *Handlers) AddSong(w http.ResponseWriter, r *http.Request) {
	var song model.Song

	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		h.Logger.Sugar.Infow("error in request handler", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(song.Song) == 0 || len(song.Group) == 0 {
		h.Logger.Sugar.Infow("error adding song", "error", errors.New("empty group name or song name"))
		http.Error(w, "empty group name or song name", http.StatusBadRequest)
		return
	}

	// params := url.Values{}
	// params.Add("group", song.Group)
	// params.Add("song", song.Song)

	// fullURL := fmt.Sprintf("%s/info?%s", h.Config.InfoService, params.Encode())

	// req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	// if err != nil {
	// 	h.Logger.Sugar.Infow("failed to create request", "error", err)
	// 	http.Error(w, "error processing request", http.StatusInternalServerError)
	// 	return
	// }

	// resp, err := h.Client.Do(req)
	// if err != nil {
	// 	h.Logger.Sugar.Infow("failed to send request", "error", err)
	// 	http.Error(w, "error processing request", http.StatusInternalServerError)
	// 	return
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	h.Logger.Sugar.Infow("received non-200 response", "status: ", resp.StatusCode)
	// 	http.Error(w, "error processing request", http.StatusInternalServerError)
	// 	return
	// }

	mockData := model.SongDetail{
		ReleaseDate: model.ReleaseDate(time.Date(2006, 07, 16, 0, 0, 0, 0, time.UTC)),
		Text:        "Ooh baby, don't you know I suffer?",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}
	mockJSON, _ := json.Marshal(mockData)

	if err := json.NewDecoder(io.NopCloser(bytes.NewReader(mockJSON))).Decode(&song); err != nil {
		// if err := json.NewDecoder(resp.Body).Decode(&song); err != nil {
		h.Logger.Sugar.Infow("failed to decode response", "error", err)
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}

	if err := h.Stor.AddSong(r.Context(), song); err != nil {
		h.Logger.Sugar.Infow("failed to add song", "error", err)
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
