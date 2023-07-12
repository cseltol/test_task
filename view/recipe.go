package view

import (
	"strings"
	"test_task/model"
)

func FormatRecipeText(recipie model.Recipe) string {
	text := RecipeTemplate
	ingridients := strings.Join(recipie.Ingridients, "; ")
	return strings.NewReplacer(
		"{name}", recipie.Name,
		"{ingridients}", ingridients,
		"{description}", recipie.Description,
	).Replace(text)
}
