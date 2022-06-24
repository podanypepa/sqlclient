package sqlclient

import (
	"database/sql"
	"fmt"
	"time"

	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

// SQLClient interface for SQL connection
type SQLClient interface {
	Query(query string, args ...any) (Rows, error)
	Ping() error
	SetMaxOpenConns(n int)
}

type client struct {
	Dsn         string
	Driver      string
	DB          *sql.DB
	ConnectedAt time.Time
}

// Open connection to mariaDB
func Open(driver string, dsn string) (SQLClient, error) {
	if driver == "" {
		return nil, fmt.Errorf("invalid driver name")
	}

	if isMocked {
		mockClient = &clientMock{}
		return mockClient, nil
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("Open err: %v", err)
	}

	sc := &client{
		Dsn:         dsn,
		Driver:      driver,
		DB:          db,
		ConnectedAt: time.Now(),
	}

	return sc, nil
}

// Query implementation of SqlClient interface
func (c *client) Query(query string, args ...any) (Rows, error) {
	res, err := c.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Query err: %v", err)
	}

	rs := &rows{
		sqlRows: res,
	}
	return rs, nil
}

// Ping implementation
func (c *client) Ping() error {
	if err := c.DB.Ping(); err != nil {
		return fmt.Errorf("Ping err %v", err)
	}
	return nil
}

// SetMaxOpenConns set the maximum possible number of open connections to
// the database from pool. Be careful — by default this value is 0, which means
// unlimited number of connections.
func (c *client) SetMaxOpenConns(n int) {
	c.DB.SetMaxOpenConns(n)
}
