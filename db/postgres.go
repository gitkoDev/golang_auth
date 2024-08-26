package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func PostgresConnection() (*pgx.Conn, error) {
	PostgresDB, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %s", err)
	}

	return PostgresDB, nil
}
