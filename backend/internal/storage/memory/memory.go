package memory

import (
	"fmt"

	"github.com/mattismoel/cookery/internal/model"
)

type memoryStorage struct{}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{}
}

func (s memoryStorage) RecipeByID(id int64) (model.Recipe, error) {
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
