package eventHandler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	responseMapper "github.com/mohsentm/event-project-test/internal/helper"
	"github.com/mohsentm/event-project-test/internal/helper/exception"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetAllEvent(w http.ResponseWriter, r *http.Request) {
	responseMapper.Success(w, getEvents())
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
	responseMapper.Success(w, newEvent)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseMapper.Exception(w, exception.BadRequest)
		return
	}

	event, err := findEvent(eventID)
	if err != nil {
		responseMapper.Exception(w, exception.NotFound)
		return
	}

	responseMapper.Success(w, event)
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
			responseMapper.Success(w, singleEvent)
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
