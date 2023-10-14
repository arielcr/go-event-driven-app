package main

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/arielcr/go-event-driven-app/event-processor-service/internal/application"
)

func main() {

	fmt.Println("- Event Processor Service -")

	router := mux.NewRouter()

	a := new(application.App)

	a.Initialize(
		router,
	)

	a.Run(":8091")

}
