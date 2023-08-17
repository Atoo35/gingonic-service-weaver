package main

import (
	"context"
	"log"

	"github.com/Atoo35/gingonic-service-weaver/api/routes"
	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	server weaver.Listener
}

func serve(ctx context.Context, app *app) error {
	router := routes.SetupRoutes()
	return router.RunListener(app.server)
}
