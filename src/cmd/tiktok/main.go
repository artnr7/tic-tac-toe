package main

import (
	"datasource"
	"net/http"
	"service_impl"
	"web"

	"go.uber.org/fx"
)

func CreateApp() fx.Option {
	return fx.Options(
		fx.Invoke(
			datasource.NewStorage,
			service_impl.NewServiceImpl,
			web.NewGameHandler,
			http.NewServeMux,
		),
		fx.Provide(RegisterRoutes),
	)
}

func RegisterRoutes(h *web.GameHandler, mtx http.ServeMux) {
	mtx.HandleFunc("POST /game/{uuid}", h.UpdateGame)
}

func main() {
	fx.New(CreateApp()).Run()
}
