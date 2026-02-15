package main

import (
	"di"

	"go.uber.org/fx"
)

func main() {
	fx.New(di.CreateApp()).Run()
}
