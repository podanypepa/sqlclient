package sqlclient

import (
	"fmt"
)

// Mock struct for testing
type Mock struct {
	Query string
	Args  []any
	Err   error
	Cols  []string
	Rows  [][]any
}

type clientMock struct {
	mocks map[string]Mock
}

var (
	isMocked   bool
	mockClient SQLClient
)

// AddMock for testing
func AddMock(mock Mock) {
	client := mockClient.(*clientMock)
	if client.mocks == nil {
		client.mocks = make(map[string]Mock)
	}
	client.mocks[mock.Query] = mock
}

// Query for mock
func (c *clientMock) Query(query string, args ...any) (Rows, error) {
	mock, exists := c.mocks[query]
	if !exists {
		return nil, fmt.Errorf("no mock for %s", query)
	}

	if mock.Err != nil {
		return nil, mock.Err
	}

	rows := rowsMock{
		Cols: mock.Cols,
		Rows: mock.Rows,
	}

	return &rows, nil
}

// Ping for mock
func (c *clientMock) Ping() error {
	return nil
}

// SetMaxOpenConns is not imppelemnted on mock struct
func (c *clientMock) SetMaxOpenConns(n int) {
}

// StartMockServer for testing
func StartMockServer() {
	isMocked = true
}

// StopMockServer for testing
func StopMockServer() {
	isMocked = false
}
