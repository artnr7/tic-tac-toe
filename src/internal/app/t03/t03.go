package t03

import (
	"io"
	"net/http"
	"os"

	"t03/internal/app/store"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type APIServer struct {
	config *Config
	logger *zerolog.Logger
	router *chi.Mux
	store  *store.Store
}

func New(config *Config) *APIServer {
	logger := zerolog.New(os.Stdout)
	router := chi.NewMux()
	return &APIServer{
		config: config,
		logger: &logger,
		router: router,
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info().Msg("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := zerolog.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	newLogger := s.logger.Level(level)

	s.logger = &newLogger

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	// здесь можно прописать что-то локально нужное
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
