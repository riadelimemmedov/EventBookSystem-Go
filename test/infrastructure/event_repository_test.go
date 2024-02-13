package infrastructure

import (
	"book_event/common/postgresql"
	"book_event/models"
	"book_event/persistence"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var eventRepository persistence.IEventRepository
var dbPool *pgxpool.Pool
var ctx context.Context

// *TestMain
func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "eventapp",
		UserName:              "postgres",
		Password:              "123321",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	eventRepository = persistence.NewEventRepository(dbPool)
	fmt.Println("Before all tests...")
	exitCode := m.Run()
	fmt.Println("After all tests...")
	os.Exit(exitCode)
}

// *setup
func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

// *clear
func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

// !TestGetAllEventsv
func TestGetAllEvents(t *testing.T) {
	setup(ctx, dbPool)
	expectedEvents := []models.Event{
		{Id: 1, Name: "Event 1", Description: "Description of Event 1", Location: "Location 1", DateTime: time.Date(2024, 2, 14, 10, 0, 0, 0, time.UTC), UserID: 1},
		{Id: 2, Name: "Event 2", Description: "Description of Event 2", Location: "Location 2", DateTime: time.Date(2024, 2, 15, 14, 30, 0, 0, time.UTC), UserID: 2},
		{Id: 3, Name: "Event 3", Description: "Description of Event 3", Location: "Location 3", DateTime: time.Date(2024, 2, 16, 18, 45, 0, 0, time.UTC), UserID: 1},
		{Id: 4, Name: "Event 4", Description: "Description of Event 4", Location: "Location 4", DateTime: time.Date(2024, 2, 17, 9, 15, 0, 0, time.UTC), UserID: 3},
		{Id: 5, Name: "Event 5", Description: "Description of Event 5", Location: "Location 5", DateTime: time.Date(2024, 2, 18, 12, 0, 0, 0, time.UTC), UserID: 2},
	}
	t.Run("GetAllEvents", func(t *testing.T) {
		actualEvents := eventRepository.GetAllEvents()
		assert.Equal(t, 5, len(actualEvents)) // there are only four events in the test database
		assert.Equal(t, expectedEvents, actualEvents)
	})
	fmt.Println("TestGetAllEvents")
	clear(ctx, dbPool)
}

// !TestAddEvent
func TestAddEvent(t *testing.T) {
	expectedEvents := []models.Event{
		{Id: 1, Name: "Event 1", Description: "Description of Event 1", Location: "Location 1", DateTime: time.Date(2024, 2, 14, 10, 0, 0, 0, time.UTC), UserID: 1},
	}

	newEvent := models.Event{
		Name: "Event 1", Description: "Description of Event 1", Location: "Location 1", DateTime: time.Date(2024, 2, 14, 10, 0, 0, 0, time.UTC), UserID: 1,
	}

	t.Run("AddEvent", func(t *testing.T) {
		err := eventRepository.AddEvent(newEvent)
		checkError(err)
		actualEvents := eventRepository.GetAllEvents()
		assert.Equal(t, 1, len(actualEvents))
		assert.Equal(t, expectedEvents, actualEvents)
	})
	clear(ctx, dbPool)
}

// ? checkError
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
