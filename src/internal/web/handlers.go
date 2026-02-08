package web

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	log.Println("create game")

	gs, err := h.s.CreateGameSession()
	if err != nil {
		http.Error(
			w,
			"can't create new game session",
			http.StatusInternalServerError,
		)
	}
	gs, err = h.s.GetGameSession(&(gs.UUID))

	if gs.CompSide == uint8(rand.Int31n(2)+1) {
		h.s.PutNextApologiseMove(gs)
	}

	dto := toDTO(gs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)

	log.Println("end create game")
}

func (h *GameHandler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	log.Println("update game")
	uuid, err := uuid.Parse(r.PathValue("uuid"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
	}

	log.Println("end parse uuid")

	// dto
	dto := NewDTO()
	dto.UUID = uuid
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON file", http.StatusBadRequest)
		return
	}

	// transformation
	gs := toDomain(dto)
	h.s.SetGameSession(gs)
	// ok, we have refreshed gs in the repo

	// business logic
	// check before
	err = h.s.GameChangeValidate(&(dto.UUID))
	if err != nil {
		log.Println("ERROR:", err)
		http.Error(w, "Game not changed", http.StatusBadRequest)
	}

	// if isn't game end
	// if h.s.IsGameEnd(&(dto.UUID)) == domain.Motive {
	// 	// prepare next move
	// 	h.s.PutNextApologiseMove(&(dto.UUID))
	// }

	// transformation
	gs, _ = h.s.GetGameSession(&(dto.UUID))
	dto = toDTO(gs)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)
	log.Println("end update game")
}
