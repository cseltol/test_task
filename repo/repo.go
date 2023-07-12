package repo

import (
	"context"
	"fmt"
	"os"
	"test_task/config"
	"test_task/model"

	"github.com/jackc/pgx/v5"
)

func InitDB() {
	c := GetConnection()
	defer c.Close(context.Background())

	c.Exec(
		context.Background(),
		"CREATE DATABASE test;",
	)
	c.Exec(
		context.Background(),
		"\\c test",
	)
	c.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS recipes (
			id serial PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			ingridients VARCHAR[] NOT NULL,
			description VARCHAR(255) NOT NULL
		);`,
	)
}

func ClearDB() {
	c := GetConnection()
	defer c.Close(context.Background())

	c.Exec(
		context.Background(),
		`DROP TABLE recipes;`,
	)

	c.Exec(
		context.Background(),
		`DROP DATABASE test;`,
	)
}

func GetConnection() *pgx.Conn {
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
