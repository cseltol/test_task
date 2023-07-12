package repo_test

import (
	"context"
	"test_task/model"
	"test_task/repo"
	"testing"
	
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetRecipeById(t *testing.T) {
	repo.ClearDB()
	repo.InitDB()
	c := repo.GetConnection()
	defer c.Close(context.Background())
	id := uuid.New().ID()
	rec_1 := &model.Recipe{
		Id: id,
		Name: "Tasty Recipe",
		Ingridients: []string{"one egg", "drop of milk"},
		Description: "Crack open the egg on the hot pan and get your drop of milk in there",
	}
	repo.CreateRecipie(c, rec_1)


	rec_2 := &model.Recipe{
		Id: id+1,
		Name: "Some Recipe",
		Ingridients: []string{"one spoon", "drop of water"},
		Description: "ooops",
	}
	repo.CreateRecipie(c, rec_2)

	actualRec, _ := repo.GetRecipieById(c, id)

	assert.Equal(t, rec_1.Id, actualRec.Id)
	assert.Equal(t, rec_1.Name, actualRec.Name)
	assert.Equal(t, rec_1.Ingridients, actualRec.Ingridients)
	assert.Equal(t, rec_1.Description, actualRec.Description)
}

func Test_GetAllRecipes(t *testing.T) {
	repo.ClearDB()
	repo.InitDB()
	c := repo.GetConnection()
	defer c.Close(context.Background())

	id := uuid.New().ID()
	rec_1 := model.Recipe{
		Id: id,
		Name: "Tasty Recipe",
		Ingridients: []string{"one egg", "drop of milk"},
		Description: "Crack open the egg on the hot pan and get your drop of milk in there",
	}
	repo.CreateRecipie(c, &rec_1)


	rec_2 := model.Recipe{
		Id: id+1,
		Name: "Some Recipe",
		Ingridients: []string{"one spoon", "drop of water"},
		Description: "ooops",
	}
	repo.CreateRecipie(c, &rec_2)

	recipes, _ := repo.GetAllRecipes(c)

	assert.Equal(t, 2, len(recipes))
}

func Test_GetFormatedRecipeById(t *testing.T) {
	repo.ClearDB()
	repo.InitDB()
	c := repo.GetConnection()
	defer c.Close(context.Background())

	id := uuid.New().ID()
	rec_1 := model.Recipe{
		Id: id,
		Name: "Tasty Recipe",
		Ingridients: []string{"one egg", "drop of milk"},
		Description: "Crack open the egg on the hot pan and get your drop of milk in there",
	}
	repo.CreateRecipie(c, &rec_1)


	rec_2 := model.Recipe{
		Id: id+1,
		Name: "Some Recipe",
		Ingridients: []string{"one spoon", "drop of water"},
		Description: "ooops",
	}
	repo.CreateRecipie(c, &rec_2)

	recipe, _ := repo.GetFormatedRecipeById(c, id)

	assert.Contains(t, recipe, "Crack open the egg on the hot pan and get your drop of milk in there")
}

func Test_EditRecipieById(t *testing.T) {
	repo.ClearDB()
	repo.InitDB()
	c := repo.GetConnection()
	defer c.Close(context.Background())

	id := uuid.New().ID()
	rec_1 := model.Recipe{
		Id: id,
		Name: "Tasty Recipe",
		Ingridients: []string{"one egg", "drop of milk"},
		Description: "Crack open the egg on the hot pan and get your drop of milk in there",
	}
	repo.CreateRecipie(c, &rec_1)


	rec_2 := model.Recipe{
		Id: id+1,
		Name: "Some Recipe",
		Ingridients: []string{"one spoon", "drop of water"},
		Description: "ooops",
	}
	repo.CreateRecipie(c, &rec_2)
	
	rec_1.Name = "Even more tasty recipe"
	rec_1.Ingridients = []string{"nothing"}
	rec_1.Description = "just order a burger"
	
	repo.EditRecipieById(c, id, rec_1)

	recipe, _ := repo.GetRecipieById(c, id)

	assert.Equal(t, "Even more tasty recipe", recipe.Name)
	assert.Equal(t, "nothing", recipe.Ingridients[0])
	assert.Equal(t, "just order a burger", recipe.Description)
}

func Test_DeleteRecipeById(t *testing.T) {
	repo.ClearDB()
	repo.InitDB()
	c := repo.GetConnection()
	defer c.Close(context.Background())

	id := uuid.New().ID()
	rec_1 := model.Recipe{
		Id: id,
		Name: "Tasty Recipe",
		Ingridients: []string{"one egg", "drop of milk"},
		Description: "Crack open the egg on the hot pan and get your drop of milk in there",
	}
	repo.CreateRecipie(c, &rec_1)


	rec_2 := model.Recipe{
		Id: id+1,
		Name: "Some Recipe",
		Ingridients: []string{"one spoon", "drop of water"},
		Description: "ooops",
	}
	repo.CreateRecipie(c, &rec_2)

	recipes, _ := repo.GetAllRecipes(c)
	err := repo.DeleteRecipeById(c, id)

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(recipes))

	err = repo.DeleteRecipeById(c, id)

	assert.NotEqual(t, nil, err)
}