package main

import (
	"go.uber.org/fx"
)

func CreateApp() fx.Option {
	return fx.Options()
}

func main() {
	fx.New(CreateApp()).Run()
}
