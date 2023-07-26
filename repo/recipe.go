package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"strings"
	"test_task/model"
	"test_task/view"
)

func GetRecipieById(c *pgx.Conn, id uint32) (model.Recipe, error) {
	var res model.Recipe
	err := c.QueryRow(
		context.Background(),
		"SELECT * FROM recipes WHERE id=$1", id,
	).Scan(&res)
	return res, err
}

func GetAllRecipes(c *pgx.Conn) ([]model.Recipe, error) {
	var res []model.Recipe
	err := c.QueryRow(
		context.Background(),
		"SELECT * FROM recipes",
	).Scan(&res)
	return res, err
}

func GetFormatedRecipeById(c *pgx.Conn, id uint32) (string, error) {
	var res model.Recipe
	err := c.QueryRow(
		context.Background(),
		"SELECT * FROM recipes WHERE id=$1", id,
	).Scan(&res)

	recipe := view.FormatRecipeText(res)

	return recipe, err
}

func EditRecipieById(c *pgx.Conn, id uint32, editedRecipe model.Recipe) error {
	_, err := c.Exec(
		context.Background(),
		"UPDATE recipes SET name=$1 ingredients=$2 description=$3 WHERE id=$4",
		editedRecipe.Name, editedRecipe.Ingredients, editedRecipe.Description, id,
	)
	return err
}

func DeleteRecipeById(c *pgx.Conn, id uint32) error {
	_, err := c.Exec(
		context.Background(),
		"DELETE FROM recipes WHERE id=$1", id,
	)
	return err
}

func FilterRecipesByIngredients(c *pgx.Conn, ingredients []string) ([]model.Recipe, error) {
	var res []model.Recipe
	ing := strings.Join(ingredients, ",")
	err := c.QueryRow(
		context.Background(),
		"SELECT * FROM recipes WHERE [$1] ~ ANY(ingredients)", ing,
	).Scan(&res)

	return res, err
}

func FilterRecipesByDuration(c *pgx.Conn) ([]model.Recipe, error) {
	var res []model.Recipe
	err := c.QueryRow(
		context.Background(),
		"SELECT * FROM recipes ORDER BY duration ASC",
	).Scan(&res)

	return res, err
}
