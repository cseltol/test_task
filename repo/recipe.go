package repo

import (
	"context"
	"test_task/model"
	"test_task/view"

	"github.com/jackc/pgx/v5"
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
		"UPDATE recipes SET name=$1 ingridients=$2 description=$3 WHERE id=$4", 
		editedRecipe.Name, editedRecipe.Ingridients, editedRecipe.Description, id,
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