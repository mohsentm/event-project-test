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

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent newEvent
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		return
	}

	_ = json.Unmarshal(reqBody, &newEvent)
	event, err := createEvent(&newEvent)
	if err != nil {
		responseMapper.Exception(w, exception.NotFound)
		return
	}

	responseMapper.Created(w, event)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseMapper.Exception(w, exception.BadRequest)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		return
	}

	var modifyEvent newEvent
	_ = json.Unmarshal(reqBody, &modifyEvent)
	event, err := updateEvent(eventID, &modifyEvent)
	if err != nil {
		responseMapper.Exception(w, exception.NotFound)
		return
	}
	responseMapper.Success(w, event)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		responseMapper.Exception(w, exception.BadRequest)
		return
	}

	deleteEvent(eventID)
	responseMapper.Deleted(w)
}
