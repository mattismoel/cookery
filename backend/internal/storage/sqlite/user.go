package sqlite

import (
	"context"
	"fmt"

	"github.com/mattismoel/cookery/internal/model"
)

func (s SQLiteStorage) UserById(ctx context.Context, id int64) (model.User, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return model.User{}, fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	query := `
  SELECT
    id,
    email,
    firstname,
    lastname,
    password_hash
  FROM
    users
  WHERE
    id = ?`

	var usr model.User

	err = tx.QueryRowContext(ctx, query, id).
		Scan(&usr.Id, &usr.Email, &usr.Firstname, &usr.Lastname, &usr.PasswordHash)

	if err != nil {
		return model.User{}, fmt.Errorf(ErrScanLayout, err)
	}

	err = tx.Commit()
	if err != nil {
		return model.User{}, fmt.Errorf(ErrCommitTxLayout, err)
	}

	return usr, nil
}

func (s SQLiteStorage) UserByEmail(ctx context.Context, email string) (model.User, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return model.User{}, fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	query := `
  SELECT
    id,
    email,
    firstname,
    lastname,
    password_hash
  FROM
    users
  WHERE
    email = ?`

	var usr model.User

	err = tx.QueryRowContext(ctx, query, email).
		Scan(&usr.Id, &usr.Email, &usr.Firstname, &usr.Lastname, &usr.PasswordHash)

	if err != nil {
		return model.User{}, fmt.Errorf(ErrScanLayout, err)
	}

	err = tx.Commit()
	if err != nil {
		return model.User{}, fmt.Errorf(ErrCommitTxLayout, err)
	}

	return usr, nil
}

func (s SQLiteStorage) AddUser(ctx context.Context, usr model.User) (int64, error) {
	fmt.Println(ctx)
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, fmt.Errorf(ErrBeginTxLayout, err)
	}

	defer tx.Rollback()

	query := `
  INSERT INTO users (
    email,
    firstname,
    lastname,
    password_hash
  )
  VALUES (?, ?, ?, ?)`

	res, err := tx.ExecContext(ctx, query, usr.Email, usr.Firstname, usr.Lastname, usr.PasswordHash)
	if err != nil {
		return -1, fmt.Errorf(ErrExecQueryLayout, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf(ErrLastInsertedIdLayout, err)
	}

	err = tx.Commit()
	if err != nil {
		return -1, fmt.Errorf(ErrCommitTxLayout, err)
	}

	return id, nil
}
