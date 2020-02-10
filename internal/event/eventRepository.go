package eventHandler

import (
	"errors"
	"fmt"

	"github.com/mohsentm/event-project-test/internal/database"
	"github.com/mohsentm/event-project-test/internal/event/model"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type newEvent struct {
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

type NotFoundExeption struct {
	Message string `json:"message"`
}

type listEvent []model.Event

func getEvents() listEvent {
	db := database.Open()
	defer database.Close()

	var events = listEvent{}
	db.Find(&events)

	return events
}

func findEvent(id int) (model.Event, error) {
	db := database.Open()
	defer database.Close()
	var event = model.Event{}

	if db.Where(&model.Event{ID: id}).First(&event).RecordNotFound() {
		fmt.Println(db.Error)
		err := errors.New("event not found")
		return model.Event{}, err
	}

	return event, nil
}

func createEvent(NewEvent *newEvent) (model.Event, error) {
	db := database.Open()
	defer database.Close()

	event := model.Event{
		Title:       NewEvent.Title,
		Description: NewEvent.Description,
	}

	if db.Create(&event).RecordNotFound() {
		fmt.Println(db.Error)
		err := errors.New("database error")
		return model.Event{}, err
	}
	return event, nil
}

func updateEvent(id int, UpdateEvent *newEvent) (model.Event, error) {
	db := database.Open()
	defer database.Close()

	var event = model.Event{}
	if db.Where(&model.Event{ID: id}).First(&event).RecordNotFound() {
		err := errors.New("event not found")
		return model.Event{}, err
	}

	event.Title = UpdateEvent.Title
	event.Description = UpdateEvent.Description
	db.Save(&event)

	return event, nil
}

func deleteEvent(id int)  {
	db := database.Open()
	defer database.Close()
	
	db.Where(&model.Event{ID: id}).Delete(&model.Event{})
}
