package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/google/uuid"
	"github.com/mattismoel/cookery/internal/filemanager"
	"github.com/mattismoel/cookery/internal/model"
	"github.com/mattismoel/cookery/internal/storage"
)

type RecipeService struct {
	storage     storage.Storage
	fileManager filemanager.FileManager
}

func NewRecipeService(storage storage.Storage, fm filemanager.FileManager) *RecipeService {
	return &RecipeService{storage: storage, fileManager: fm}
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

func (s RecipeService) SetBannerImage(ctx context.Context, id int64, fh *multipart.FileHeader) (string, error) {
	ext := path.Ext(fh.Filename)
	bannerName := fmt.Sprintf("recipes/banner-%s%s", uuid.NewString(), ext)
	bannerFile, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("could not open banner file: %v", err)
	}

	defer bannerFile.Close()

	bannerUrl, err := s.fileManager.Save(ctx, bannerFile, bannerName)
	if err != nil {
		return "", fmt.Errorf("could not save banner image: %v", err)
	}

	err = s.storage.SetBannerUrl(ctx, id, bannerUrl)
	if err != nil {
		return "", fmt.Errorf("could not set banner url: %v", err)
	}

	return bannerUrl, nil
}
