package sqlite

import (
	"context"
	"fmt"
)

func (s SQLiteStorage) createUsersTable(ctx context.Context) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	query := `
  CREATE TABLE IF NOT EXISTS users (
    id                INTEGER     PRIMARY KEY       NOT NULL,
    email             TEXT                          NOT NULL,
    firstname         TEXT                          NOT NULL,
    lastname          TEXT                          NOT NULL,
    password_hash     TEXT                          NOT NULL
  )`

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf(ErrExecQueryLayout, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(ErrCommitTxLayout, err)
	}

	return nil
}

func (s SQLiteStorage) createRecipesTable(ctx context.Context) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	query := `
  CREATE TABLE IF NOT EXISTS recipes (
    id                INTEGER       PRIMARY KEY     NOT NULL,
    title             TEXT                          NOT NULL,
    banner_url        TEXT,
    cook_minutes      INTEGER                       NOT NULL,
    total_minutes     INTEGER                       NOT NULL,
    author_id         INTEGER                       NOT NULL,

    FOREIGN KEY (author_id) REFERENCES users(id)
  )`

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf(ErrInitTableLayout, "recipes", err)
	}

	query = `
  CREATE TABLE IF NOT EXISTS subrecipes (
    id            INTEGER     PRIMARY KEY     NOT NULL,
    title         TEXT                        NOT NULL,
    recipe_id     INTEGER                     NOT NULL,

    FOREIGN KEY (recipe_id) REFERENCES recipes(id)
  )`

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf(ErrInitTableLayout, "subrecipes", err)
	}

	query = `
  CREATE TABLE IF NOT EXISTS ingredients (
    id                INTEGER     PRIMARY KEY     NOT NULL,
    name              TEXT                        NOT NULL,
    unit              TEXT,
    note              TEXT,

    subrecipe_id      INTEGER                     NOT NULL,

    FOREIGN KEY (subrecipe_id) REFERENCES subrecipe(id)
  )`

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf(ErrInitTableLayout, "ingredients", err)
	}

	query = `
  CREATE TABLE IF NOT EXISTS instructions (
  id                INTEGER     PRIMARY KEY     NOT NULL,
  text              TEXT                        NOT NULL,
  image_url         TEXT,
  subrecipe_id      INTEGER                     NOT NULL,

  FOREIGN KEY (subrecipe_id) REFERENCES subrecipe(id)
  )`

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf(ErrInitTableLayout, "instructions", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(ErrCommitTxLayout, err)
	}

	return nil
}
