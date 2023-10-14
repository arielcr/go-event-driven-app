package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arielcr/go-event-driven-app/event-processor-service/internal/events"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(r *mux.Router) {
	a.Router = r

	a.InitializeRoutes()
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/event", a.storeEvent).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Printf("Service Initialized at port %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) storeEvent(w http.ResponseWriter, r *http.Request) {
	var e events.Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	respondWithJSON(w, http.StatusCreated, e)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
