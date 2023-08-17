package main

import (
	"context"
	"log"

	"github.com/Atoo35/gingonic-service-weaver/taskservice"
	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), taskservice.Serve); err != nil {
		log.Fatal(err)
	}
}
