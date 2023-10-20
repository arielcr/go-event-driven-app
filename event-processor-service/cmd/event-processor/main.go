package main

import (
	"github.com/arielcr/go-event-driven-app/event-processor-service/internal/application"
)

func main() {
	a := application.NewService()

	a.Initialize()

	a.Run(":8091")

}
