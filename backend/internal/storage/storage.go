package storage

import "github.com/mattismoel/cookery/internal/model"

type Storage interface {
	RecipeByID(int64) (model.Recipe, error)
	AllRecipes() ([]model.Recipe, error)
}
