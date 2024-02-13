package infrastructure

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
)

var INSERT_EVENT = `
	INSERT INTO event (name, description, location, dateTime, user_id)
	VALUES
		('Event 1', 'Description of Event 1', 'Location 1', '2024-02-14 10:00:00', 1),
		('Event 2', 'Description of Event 2', 'Location 2', '2024-02-15 14:30:00', 2),
		('Event 3', 'Description of Event 3', 'Location 3', '2024-02-16 18:45:00', 1),
		('Event 4', 'Description of Event 4', 'Location 4', '2024-02-17 09:15:00', 3),
		('Event 5', 'Description of Event 5', 'Location 5', '2024-02-18 12:00:00', 2);
`

// ! TestDataInitialize
func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertEventsResult, insertEventsErr := dbPool.Exec(ctx, INSERT_EVENT)
	if insertEventsErr != nil {
		log.Error("Occur problem when save to database")
	} else {
		log.Info(fmt.Sprintf("Events data created with %d rows,how many rows affected...", insertEventsResult.RowsAffected()))
	}
}
