package memory

import (
	"fmt"

	"github.com/mattismoel/cookery/internal/model"
)

var recipes []model.Recipe = []model.Recipe{
	{
		Id:           0,
		Title:        "Italian Chicken Pasta",
		Author:       "John Doe",
		BannerUrl:    "https://images.immediate.co.uk/production/volatile/sites/30/2013/05/2022-09-23GFOWEBFamilyRecipes-OnePotGarlicChicken05875preview-d8a4a42.jpg",
		CookMinutes:  15,
		TotalMinutes: 45,
		SubRecipes: []model.SubRecipe{
			{
				Title: "Pasta",
				Ingredients: []model.Ingredient{
					{Name: "pasta", Amount: 250, Unit: "g"},
					{Name: "garlic", Amount: 2, Unit: "cloves"},
					{Name: "cream", Amount: 200, Unit: "mL"},
					{Name: "parsley", Amount: 2, Unit: "handful"},
					{Name: "chicken breast", Amount: 2},
				},
				Instructions: []model.Instruction{
					{Text: "Begin boiling water for the pasta. Preheat a pan with a generous amount of olive oil, and cook the chopped garlic."},
					{Text: "When the garlic is browned add the chicken breast to the pan and cook till brown on one side. Once the side is brown flip it."},
				},
			},
			{
				Title: "Caesar Salad",
				Ingredients: []model.Ingredient{
					{Name: "ruccola", Amount: 100, Unit: "g"},
					{Name: "tomatoes", Amount: 8},
					{Name: "vinegar", Amount: 10, Unit: "mL", Note: "apple cider"},
					{Name: "extra virgin olive oil", Amount: 10, Unit: "mL"},
				},
				Instructions: []model.Instruction{
					{Text: "Chop the tomatoes in half."},
					{Text: "Mix ruccola and tomatoes in a bowl."},
				},
			},
		},
	},
}

func (s memoryStorage) RecipeById(id int64) (model.Recipe, error) {
	for _, recipe := range recipes {
		if recipe.Id == id {
			return recipe, nil
		}
	}

	return model.Recipe{}, fmt.Errorf("could not find recipe with id %d", id)
}

func (s memoryStorage) AllRecipes() ([]model.Recipe, error) {
	return recipes, nil
}
