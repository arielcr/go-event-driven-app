package application

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/arielcr/go-event-driven-app/event-processor-service/internal/events"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	Router *mux.Router
	Log    *log.Logger
}

func NewService() *Service {
	s := Service{
		mux.NewRouter(),
		log.New(),
	}
	return &s
}

func (s *Service) Initialize() {
	s.InitializeRoutes()
	s.InitializeLogger()
	s.Log.Println("Event Procesor Service started")
}

func (s *Service) InitializeRoutes() {
	s.Router.HandleFunc("/event", s.storeEvent).Methods("POST")
}

func (s *Service) InitializeLogger() {
	s.Log.SetOutput(os.Stdout)

	s.Log.SetLevel(log.DebugLevel)

	s.Log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func (s *Service) Run(addr string) {
	s.Log.Printf("Service Initialized at port %s\n", addr)
	s.Log.Fatal(http.ListenAndServe(addr, s.Router))
}

func (s *Service) storeEvent(w http.ResponseWriter, r *http.Request) {
	var e events.Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	s.Log.WithFields(log.Fields{
		"payload": e,
	}).Info("Storing Event in DynamoDB")

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
