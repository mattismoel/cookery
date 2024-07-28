package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrBeginTxLayout        = "could not begin transaction: %v"
	ErrCommitTxLayout       = "could not commit transaction: %v"
	ErrExecQueryLayout      = "could not execute query: %v"
	ErrInitTableLayout      = "could not initialise table :%q: %v"
	ErrScanLayout           = "could not scan into struct: %v"
	ErrLastInsertedIdLayout = "could not get last inserted id: %v"
)

type SQLiteStorage struct {
	dbPath string
	db     *sql.DB
}

func New(dbPath string) *SQLiteStorage {
	return &SQLiteStorage{dbPath: dbPath}
}

func (s *SQLiteStorage) Start() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("sqlite3", s.dbPath)
	if err != nil {
		return fmt.Errorf("could not open sqlite database file: %v", err)
	}

	s.db = db

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("could not ping sqlite database: %v", err)
	}

	err = s.createTables(ctx)
	if err != nil {
		return fmt.Errorf("could not create tables: %v", err)
	}

	return nil
}

func (s SQLiteStorage) createTables(ctx context.Context) error {
	err := s.createUsersTable(ctx)
	if err != nil {
		return fmt.Errorf("could not create users table: %v", err)
	}

	return nil
}
