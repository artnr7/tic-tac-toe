package web

import (
	"domain"
	"encoding/json"
	"log"
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
	defer log.Println("end create game")

	gs, err := h.s.CreateGameSession()
	if err != nil {
		http.Error(
			w,
			"can't create new game session",
			http.StatusInternalServerError,
		)
	}
	h.s.MakeNextMove(gs)

	err = h.s.PutGameSession(gs)
	if err != nil {
		http.Error(
			w,
			"can't create new game session: db error",
			http.StatusInternalServerError,
		)
	}

	dto := toDTO(gs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)
}

func (h *GameHandler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	log.Println("update game")
	defer log.Println("end update game")

	// parsing
	uuid, err := uuid.Parse(r.PathValue("uuid"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
	}

	// dto
	dto := NewDTO()
	dto.UUID = uuid
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		// fmt.Println(err)
		http.Error(w, "Invalid JSON file", http.StatusBadRequest)
		return
	}

	gs := toDomain(dto)

	// business logic ----
	err = h.s.GameChangeValidate(gs, &(gs.UUID))
	if err != nil {
		// log.Println("ERROR:", err)
		http.Error(w, "Game not changed", http.StatusBadRequest)
	}

	// game status check -----
	h.s.IsGameEnd(gs)
	if gs.CompStatus == domain.Motive {
		h.s.MakeNextMove(gs)
	}

	dto = toDTO(gs)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)
}
