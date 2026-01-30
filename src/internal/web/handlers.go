package web

import (
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
	uuid, err := uuid.Parse(r.PathValue("uuid"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
	}

	// dto
	var dto dto
	dto.uuid = uuid
	if json.NewDecoder(r.Body).Decode(dto) != nil {
		http.Error(w, "Invalid JSON file", http.StatusBadRequest)
		return
	}

	// transformation
	gs := toDomain(&dto)
	h.s.SetGameSession(gs)

	// business logic
	// check before
	err = h.s.GameChangeValidate()
	if err != nil {
		http.Error(w, "Game not changed", http.StatusBadRequest)
	}

	end := h.s.IsGameEnd()
	if end {
	} else {
		// prepare next move
		h.s.PutNextApologiseMove()
	}

	// transformation
	gs1 := h.s.GetGameSession(dto.uuid)
	dto1 := toDTO(&gs1)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto1)
}
