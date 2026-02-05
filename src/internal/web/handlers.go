package web

import (
	"encoding/json"
	"html/template"
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

func (h *GameHandler) Root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/index.html"))

	tmpl.Execute(w, nil)
}

func (h *GameHandler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(r.PathValue("uuid"))
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
	}

	// dto
	var dto *dto
	dto.UUID = uuid
	if json.NewDecoder(r.Body).Decode(dto) != nil {
		http.Error(w, "Invalid JSON file", http.StatusBadRequest)
		return
	}

	// check this game is existed
	// if it existed then we send it to user
	if _, err := h.s.GetGameSession(&uuid); err != nil {
		gs := h.s.CreateGameSession(&uuid)
		dto := toDTO(gs)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto)
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
}
