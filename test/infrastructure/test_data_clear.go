package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// !TruncateTestData
func TruncateTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	_, truncateResultErr := dbPool.Exec(ctx, "TRUNCATE event RESTART IDENTITY")
	if truncateResultErr != nil {
		log.Error("Not event table exists on database")
	} else {
		log.Info("Event table truncated")
	}
}
