package main

import (
	"datasource"
	"net/http"
	"service"
	"service_impl"
	"web"

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
	mtx.HandleFunc("POST /game/{uuid}/", h.UpdateGame)
	mtx.HandleFunc("GET /", h.Root)
	// mtx.Handle("GET /static/", http.FileServer(http.Dir("./static")))
	// mtx.Handle(
	// 	"GET /static/",
	// 	http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	// )
	// mtx.Handle("/", http.FileServer(http.Dir("./web/")))
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", mtx)
}

func main() {
	// storage := datasource.NewMap()
	// service := service_impl.NewServiceImpl(storage)
	// web := web.NewGameHandler(service)
	//
	// mtx := http.NewServeMux()
	// mtx.HandleFunc("POST /game/{uuid}/", web.UpdateGame)
	// http.ListenAndServe(":8080", mtx)

	fx.New(CreateApp()).Run()
}
