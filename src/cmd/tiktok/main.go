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
			datasource.NewMap,
			service_impl.NewServiceImpl,
			web.NewGameHandler,
		),
		fx.Provide(RegisterRoutes),
	)
}

func RegisterRoutes(h *web.GameHandler) {
	mtx := http.NewServeMux()
	mtx.HandleFunc("POST /game/{uuid}", h.UpdateGame)
	http.ListenAndServe(":8080", mtx)
}

func main() {
	fx.New(CreateApp()).Run()
}
