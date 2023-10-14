package events

type Event struct {
	EventType string    `json:"event_type"`
	EventData EventData `json:"event_data"`
}

type EventData struct {
	Client  string `json:"client"`
	Product string `json:"product"`
}
