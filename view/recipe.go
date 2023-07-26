package view

import (
	"strconv"
	"strings"
	"test_task/model"
)

func FormatRecipeText(recipe model.Recipe) string {
	text := RecipeTemplate
	ingredients := strings.Join(recipe.Ingredients, "; ")
	steps := formatStepsText(recipe.Steps)
	return strings.NewReplacer(
		"{name}", recipe.Name,
		"{ingredients}", ingredients,
		"{description}", recipe.Description,
		"{steps}", steps,
	).Replace(text)
}

func formatStepsText(steps []model.Step) string {
	t := "{step_num}. {step_description}. {step_duration} min."
	for i, s := range steps {
		stepNum := strconv.Itoa(i) + ". "
		stepTime := strconv.Itoa(
			int(s.StepDuration.Minutes()),
		)
		strings.NewReplacer(
			"{step_num}", stepNum,
			"{step_description}", s.StepDescription,
			"{step_duration}", stepTime,
		).Replace(t)
	}
	return t
}
