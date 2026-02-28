package di

import (
	"net/http"

	"tictactoe/internal/datasource"
	"tictactoe/internal/service"
	"tictactoe/internal/service_impl"
	"tictactoe/internal/web"

	"go.uber.org/fx"
)

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(datasource.NewMap, fx.As(new(service.Repository))),
			fx.Annotate(service_impl.NewServiceImpl,
				fx.As(new(service.Service))),
			http.NewServeMux,
			web.NewGameHandler,
		),
		fx.Invoke(RegisterRoutes),
	)
}

func RegisterRoutes(h *web.GameHandler, mtx *http.ServeMux) {
	mtx.HandleFunc("POST /game/{uuid}", h.UpdateGame)
	mtx.HandleFunc("POST /game/create_game", h.CreateGame)
	mtx.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", mtx)
}
