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

	// Id           int64  `json:"id"`
	// Email        string `json:"email"`
	// Firstname    string `json:"firstname"`
	// Lastname     string `json:"lastname"`
	// PasswordHash string `json:"passwordHash,omitempty"`

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
