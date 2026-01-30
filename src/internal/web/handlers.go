package web

import (
	"domain"
	"encoding/json"
	"net/http"
	"service"

	"github.com/google/uuid"
)

type GameHandler struct {
	s service.Service
}

func NewGameHandler(s service.Service) *GameHandler {
	return &GameHandler{s: s}
}

func (h *GameHandler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(r.PathValue("UUID"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
	}

	var g domain.GameSession
	if json.NewDecoder(r.Body).Decode(g) != nil {
		http.Error(w, "Invalid JSON file", http.StatusBadRequest)
		return
	}

	// business logic
	// check before
	err = g.GameChangeValidate()
	if err != nil {
		http.Error(w, "Game not changed", http.StatusBadRequest)
	}

	// saving validated game
	h.r.SaveGame(uuid, g)

	// prepare next move

	g.PutNextApologiseMove()

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode()
}
