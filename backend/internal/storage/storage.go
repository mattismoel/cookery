package storage

import (
	"context"

	"github.com/mattismoel/cookery/internal/model"
)

type Storage interface {
	UserById(context.Context, int64) (model.User, error)
	UserByEmail(context.Context, string) (model.User, error)
	AddUser(context.Context, model.User) (int64, error)

	RecipeById(context.Context, int64) (model.Recipe, error)
	AllRecipes(context.Context) ([]model.Recipe, error)
	InsertRecipe(context.Context, model.Recipe) (int64, error)
}
