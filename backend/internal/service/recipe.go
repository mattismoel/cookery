package service

import (
	"context"
	"fmt"

	"github.com/mattismoel/cookery/internal/model"
	"github.com/mattismoel/cookery/internal/storage"
)

type RecipeService struct {
	storage storage.Storage
}

func NewRecipeService(storage storage.Storage) *RecipeService {
	return &RecipeService{storage: storage}
}

func (s RecipeService) ById(ctx context.Context, id int64) (model.Recipe, error) {
	rcp, err := s.storage.RecipeById(ctx, id)
	if err != nil {
		return model.Recipe{}, fmt.Errorf("could not get recipe with id %d from storage: %v", id, err)
	}

	return rcp, nil
}

func (s RecipeService) Add(ctx context.Context, rcp model.Recipe) (int64, error) {
	id, err := s.storage.InsertRecipe(ctx, rcp)
	if err != nil {
		return -1, fmt.Errorf("could not insert recipe into storage: %v", err)
	}

	return id, nil
}
