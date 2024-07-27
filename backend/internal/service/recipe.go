package service

import (
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

func (s RecipeService) ById(id int64) (model.Recipe, error) {
	rcp, err := s.storage.RecipeById(id)
	if err != nil {
		return model.Recipe{}, fmt.Errorf("could not get recipe with id %d from storage: %v", id, err)
	}

	return rcp, nil
}
