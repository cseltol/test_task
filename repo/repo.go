package repo

import (
	"context"
	"fmt"
	"os"
	"test_task/config"
	"test_task/model"

	"github.com/jackc/pgx/v5"
)

func GetConnection() *pgx.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), config.DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func CreateRecipie(c *pgx.Conn, recipe *model.Recipe) error {
	return c.QueryRow(context.Background(),
		"INSERT INTO recipes (name, ingridients, description) VALUES ($1, $2, $3) RETURNING id",
		recipe.Name, recipe.Ingridients, recipe.Description,
	).Scan(&recipe.Id)
}