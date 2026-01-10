package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DatabaseStart() {
	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		dbHost = "localhost"
	}
	
	connString := fmt.Sprintf("postgres://marcus:12345678@%s:5432/atlas-db", dbHost)
	conn, err := pgx.Connect(context.Background(), connString)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
}
