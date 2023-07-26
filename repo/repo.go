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
		`CREATE TYPE IF NOT EXISTS step AS (
			stepDuration interval ,
			stepDescription VARCHAR[] NOT NULL
		);`,
	)
	c.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS recipes (
			id serial PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			ingredients VARCHAR[] NOT NULL,
			description VARCHAR(255) NOT NULL,
		   	steps step[] NOT NULL
		);`,
	)
	c.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			email VARCHAR(50) NOT NULL,
			password VARCHAR(50) NOT NULL,
			encryptedPassword VARCHAR(255)
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
		"INSERT INTO recipes (createdById, name, ingredients, description, steps) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		recipe.CreatedById, recipe.Name, recipe.Ingredients, recipe.Description, recipe.Steps,
	).Scan(&recipe.Id)
}
