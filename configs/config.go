package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DatabaseStart() {
	conn, err := pgx.Connect(context.Background(), "postgres://marcus:12345678@localhost:5432/atlas-db")
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
}
