package sqlclient

import (
	"database/sql"
	"fmt"
)

type rows struct {
	sqlRows *sql.Rows
}

// Rows is sql.Rows
type Rows interface {
	Next() bool
	Close() error
	Scan(dest ...any) error
}

// Next is true if in sqlRows is another row
func (r *rows) Next() bool {
	return r.sqlRows.Next()
}

// Close the connection
func (r *rows) Close() error {
	if err := r.sqlRows.Close(); err != nil {
		return fmt.Errorf("Close err: %v", err)
	}

	return nil
}

// Scan values to dest arguments
func (r *rows) Scan(dest ...any) error {
	if err := r.sqlRows.Scan(dest...); err != nil {
		return fmt.Errorf("Scan err: %v", err)
	}

	return nil
}
