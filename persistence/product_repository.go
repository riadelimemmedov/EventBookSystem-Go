package persistence

import (
	"book_event/models"
	"book_event/persistence/common"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IEventRepository interface {
	GetAllEvents() []models.Event
	AddEvent(event models.Event) error
	GetEventById(eventId int64) (models.Event, error)
}

type EventRepository struct {
	dbPool *pgxpool.Pool
}

// *NewEventRepository
func NewEventRepository(dbPool *pgxpool.Pool) IEventRepository {
	return &EventRepository{dbPool: dbPool}
}

// !GetAllEvents
func (eventRepository *EventRepository) GetAllEvents() []models.Event {
	ctx := context.Background()

	eventRows, err := eventRepository.dbPool.Query(ctx, "SELECT * FROM event")
	if err != nil {
		log.Error("Error while getting products", err)
	}
	return extractProductsFromRows(eventRows)
}

// !AddEvent
func (eventRepository *EventRepository) AddEvent(event models.Event) error {
	ctx := context.Background()

	insertProductSql := `INSERT INTO event (name,description,location,dateTime,user_id) VALUES($1,$2,$3,$4,$5)`

	addNewProduct, err := eventRepository.dbPool.Exec(ctx, insertProductSql, event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		log.Error("Failed to add new event", err)
		return err
	} else {
		log.Info(fmt.Printf("Product added to database successfully: %s", addNewProduct))
	}
	return nil
}

// !GetEventById
func (eventRepository *EventRepository) GetEventById(eventId int64) (models.Event, error) {
	ctx := context.Background()

	fmt.Println("Bunedir ala bele ", eventId)
	getEventById := `SELECT * FROM event WHERE id=$1`

	queryRow := eventRepository.dbPool.QueryRow(ctx, getEventById, eventId)

	var id int64
	var name string
	var description string
	var location string
	var dateTime time.Time
	var userId int64

	scanErr := queryRow.Scan(&id, &name, &description, &location, &dateTime, &userId)

	if scanErr != nil && scanErr.Error() == common.NOT_FOUND {
		return models.Event{}, errors.New(fmt.Sprintf("Product not found with id: %d", eventId))
	}
	if scanErr != nil {
		return models.Event{}, errors.New(fmt.Sprintf("Database error when getting product by id: %v", eventId))
	}

	return models.Event{
		Id:          id,
		Name:        name,
		Description: description,
		Location:    location,
		DateTime:    dateTime,
		UserID:      userId,
	}, nil

}

// *extractProductsFromRows
func extractProductsFromRows(productRows pgx.Rows) []models.Event {
	var events = []models.Event{}
	var id int64
	var name string
	var description string
	var location string
	var dateTime time.Time
	var userId int64

	for productRows.Next() {
		productRows.Scan(&id, &name, &description, &location, &dateTime, &userId)
		events = append(events, models.Event{
			Id:          id,
			Name:        name,
			Description: description,
			Location:    location,
			DateTime:    dateTime,
			UserID:      userId,
		})
	}
	return events
}
