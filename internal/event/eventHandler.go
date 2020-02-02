package eventHandler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	responseMapper "github.com/mohsentm/event-project-test/internal/helper"
	"github.com/mohsentm/event-project-test/internal/database"
	"github.com/mohsentm/event-project-test/internal/event/model"
	"io/ioutil"
	"net/http"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AllEvents []event

var events = AllEvents{
	{
		ID:          "1",
		Title:       "test title",
		Description: "test description for this",
	},
}

func GetAllEvent(w http.ResponseWriter, r *http.Request) {
	db := database.Open()
	defer database.Close()

	var result model.Event
	db.Find(&result)

	responseMapper.ResponseHandler(w, result)
	// responseMapper.ResponseHandler(w, events)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	_ = json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)
	responseMapper.ResponseHandler(w, newEvent)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			responseMapper.ResponseHandler(w, singleEvent)
		}
	}
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	var updateEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	_ = json.Unmarshal(reqBody, &updateEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updateEvent.Title
			singleEvent.Description = updateEvent.Description
			events = append(events[:i], singleEvent)
			responseMapper.ResponseHandler(w, singleEvent)
		}
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	var updateEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	_ = json.Unmarshal(reqBody, &updateEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			_, _ = fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}
