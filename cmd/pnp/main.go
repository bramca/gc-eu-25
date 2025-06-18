package main

import (
	"github.com/ronna-s/gc-eu-25/pkg/pnp"
	engine "github.com/ronna-s/gc-eu-25/pkg/pnp/engine/tview"
)

func main() {
	var pm pnp.ProductManager
	app := pnp.New(&pm)
	app.Run(engine.New())
}
