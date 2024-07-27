package storage

import "github.com/mattismoel/cookery/internal/model"

type Storage interface {
	UserById(int64) (model.User, error)
	AddUser(model.User) (int64, error)

	RecipeById(int64) (model.Recipe, error)
	AllRecipes() ([]model.Recipe, error)
}
