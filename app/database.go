package app

import (
	"context"
	"fmt"
	"os"
	"tobialbertino/portfolio-be/config"

	"github.com/jackc/pgx/v5"
)

func NewDB(cfg *config.Config) *pgx.Conn {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s`,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return conn
}
