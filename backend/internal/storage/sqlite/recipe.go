package sqlite

import (
	"context"
	"fmt"

	"github.com/mattismoel/cookery/internal/model"
)

func (s SQLiteStorage) RecipeById(ctx context.Context, id int64) (model.Recipe, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return model.Recipe{}, fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	// TODO: Implement SubRecipes   []SubRecipe `json:"subRecipes"`
	query := `
  SELECT
    id,
    title,
    author,
    banner_url,
    cook_minutes,
    total_minutes
  FROM
    recipes
  WHERE
    id = ?`

	var rcp model.Recipe
	err = tx.QueryRowContext(ctx, query, id).
		Scan(
			&rcp.Id,
			&rcp.Title,
			&rcp.Author,
			&rcp.BannerUrl,
			&rcp.CookMinutes,
			&rcp.TotalMinutes,
		)

	if err != nil {
		return model.Recipe{}, fmt.Errorf(ErrScanLayout, err)
	}

	err = tx.Commit()
	if err != nil {
		return model.Recipe{}, fmt.Errorf(ErrCommitTxLayout, err)
	}

	return rcp, nil
}

func (s SQLiteStorage) AllRecipes(ctx context.Context) ([]model.Recipe, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	// TODO: Implement SubRecipes   []SubRecipe `json:"subRecipes"`
	query := `
  SELECT
    id,
    title,
    author,
    banner_url,
    cook_minutes,
    total_minutes
  FROM
    recipes`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(ErrExecQueryLayout, err)
	}

	rcps := []model.Recipe{}

	for rows.Next() {
		var rcp model.Recipe
		err = rows.Scan(
			&rcp.Id,
			&rcp.Title,
			&rcp.Author,
			&rcp.BannerUrl,
			&rcp.CookMinutes,
			&rcp.TotalMinutes,
		)

		if err != nil {
			return nil, fmt.Errorf(ErrScanLayout, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf(ErrCommitTxLayout, err)
	}

	return rcps, nil
}
