package filemanager

import (
	"context"
	"io"
)

// A base interface for a FileManager, responsible for simple read, write,
// delete functionality.
type FileManager interface {
	// Saves data to a specified path.
	// The function returns the actual absolute path where the file was stored.
	Save(context.Context, io.Reader, string) (string, error)

	// Deletes a file at the given path.
	Delete(context.Context, string) error

	// Gets a file at the given path.
	Get(context.Context, string) (io.Reader, error)
}
