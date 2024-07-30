package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mattismoel/cookery/internal/model"
)

func (s SQLiteStorage) RecipeById(ctx context.Context, id int64) (model.Recipe, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return model.Recipe{}, fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	query := `
  SELECT
    r.id,
    r.title,
    r.cook_minutes,
    r.total_minutes
  FROM
    recipes r
  WHERE
    r.id = ?`

	var rcp model.Recipe
	err = tx.QueryRowContext(ctx, query, id).
		Scan(
			&rcp.Id,
			&rcp.Title,
			// &rcp.Author,
			// &rcp.BannerUrl,
			&rcp.CookMinutes,
			&rcp.TotalMinutes,
		)

	if err != nil {
		return model.Recipe{}, fmt.Errorf(ErrScanLayout, err)
	}

	rcp.SubRecipes, err = s.SubRecipes(tx, id)
	if err != nil {
		return model.Recipe{}, fmt.Errorf("could not get sub recipes: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return model.Recipe{}, fmt.Errorf(ErrCommitTxLayout, err)
	}

	return rcp, nil
}

func (s SQLiteStorage) SubRecipes(tx *sql.Tx, recipeId int64) ([]model.SubRecipe, error) {
	query := "SELECT id, title FROM subrecipes WHERE recipe_id = ?"

	fmt.Println("rcpid", recipeId)
	rows, err := tx.Query(query, recipeId)
	if err != nil {
		return nil, fmt.Errorf(ErrExecQueryLayout, err)
	}

	subRcps := []model.SubRecipe{}
	for rows.Next() {
		var subRcp model.SubRecipe
		var subRcpId int64

		err := rows.Scan(&subRcpId, &subRcp.Title)
		if err != nil {
			return nil, fmt.Errorf(ErrScanLayout, err)
		}

		subRcp.Ingredients, err = s.Ingredients(tx, subRcpId)
		if err != nil {
			return nil, fmt.Errorf("could not get ingredients: %v", err)
		}

		subRcp.Instructions, err = s.Instructions(tx, subRcpId)
		if err != nil {
			return nil, fmt.Errorf("could not get instructions: %v", err)
		}

		subRcps = append(subRcps, subRcp)
	}

	return subRcps, nil
}

func (s SQLiteStorage) Ingredients(tx *sql.Tx, subRcpId int64) ([]model.Ingredient, error) {
	query := "SELECT name, amount, unit FROM ingredients where subrecipe_id = ?"

	rows, err := tx.Query(query, subRcpId)
	if err != nil {
		return nil, fmt.Errorf(ErrExecQueryLayout, err)
	}

	ingredients := []model.Ingredient{}
	for rows.Next() {
		var ingredient model.Ingredient
		err = rows.Scan(&ingredient.Name, &ingredient.Amount, &ingredient.Unit)
		if err != nil {
			return nil, fmt.Errorf(ErrScanLayout, err)
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (s SQLiteStorage) Instructions(tx *sql.Tx, subRcpId int64) ([]model.Instruction, error) {
	query := "SELECT text, image_url FROM instructions WHERE subrecipe_id = ?"

	rows, err := tx.Query(query, subRcpId)
	if err != nil {
		return nil, fmt.Errorf(ErrExecQueryLayout, err)
	}

	instructions := []model.Instruction{}
	for rows.Next() {
		var instruction model.Instruction
		err = rows.Scan(&instruction.Text, &instruction.ImageUrl)
		if err != nil {
			return nil, fmt.Errorf(ErrScanLayout, err)
		}

		instructions = append(instructions, instruction)
	}

	return instructions, nil
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

// TODO: Implement author. Temporarily author ID = 0.
func (s SQLiteStorage) InsertRecipe(ctx context.Context, rcp model.Recipe) (int64, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	query := `
  INSERT INTO recipes (
    title,
    author_id,
    banner_url,
    cook_minutes,
    total_minutes
  )
  VALUES (?, ?, ?, ?, ?)`

	res, err := tx.ExecContext(
		ctx,
		query,
		rcp.Title,
		0,
		rcp.BannerUrl,
		rcp.CookMinutes,
		rcp.TotalMinutes,
	)

	if err != nil {
		return -1, fmt.Errorf(ErrExecQueryLayout, err)
	}

	recipeId, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf(ErrLastInsertedIdLayout, err)
	}

	insertSubRecipeQuery := `
  INSERT INTO subrecipes (title, recipe_id) VALUES (?, ?)`

	for _, subRecipe := range rcp.SubRecipes {
		res, err := tx.ExecContext(ctx, insertSubRecipeQuery, subRecipe.Title, recipeId)
		if err != nil {
			return -1, fmt.Errorf(ErrExecQueryLayout, err)
		}

		subRecipeId, err := res.LastInsertId()
		if err != nil {
			return -1, fmt.Errorf(ErrLastInsertedIdLayout, err)
		}

		err = s.AddIngredients(tx, subRecipe.Ingredients, subRecipeId)
		if err != nil {
			return -1, fmt.Errorf("could not add ingredients: %v", err)
		}

		err = s.AddInstructions(tx, subRecipe.Instructions, subRecipeId)
		if err != nil {
			return -1, fmt.Errorf("could not add instructions: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return -1, fmt.Errorf(ErrCommitTxLayout, err)
	}

	return recipeId, nil
}

func (s SQLiteStorage) AddIngredients(tx *sql.Tx, ingredients []model.Ingredient, subRcpId int64) error {
	query := `
    INSERT INTO ingredients (
      name,
      amount,
      unit,
      subrecipe_id
    )
    VALUES (?, ?, ?, ?)`

	for _, ingredient := range ingredients {
		_, err := tx.Exec(
			query,
			ingredient.Name,
			ingredient.Amount,
			ingredient.Unit,
			subRcpId,
		)

		if err != nil {
			return fmt.Errorf(ErrExecQueryLayout, err)
		}
	}

	return nil
}

func (s SQLiteStorage) AddInstructions(tx *sql.Tx, instructions []model.Instruction, subRcpId int64) error {
	query := `
    INSERT INTO instructions (
      text,
      image_url,
      subrecipe_id
    )
    VALUES (?, ?, ?)`

	for _, instruction := range instructions {
		_, err := tx.Exec(
			query,
			instruction.Text,
			instruction.ImageUrl,
			subRcpId,
		)

		if err != nil {
			return fmt.Errorf(ErrExecQueryLayout, err)
		}
	}

	return nil
}
