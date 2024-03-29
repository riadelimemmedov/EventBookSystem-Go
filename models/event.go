package models

import "time"

//*Event
type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

//!Save
func (e Event) Save() {
	//Add to database
	events = append(events, e)
}

//!GetAllEvents
func GetAllEvents() []Event {
	return events
}
