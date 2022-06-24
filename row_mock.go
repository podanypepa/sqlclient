package sqlclient

import (
	"fmt"
	"reflect"
)

type rowsMock struct {
	Cols  []string
	Rows  [][]any
	index int
}

// Next for mocked Rpws
func (m *rowsMock) Next() bool {
	return m.index < len(m.Rows)
}

// Close mocked connection
func (m *rowsMock) Close() error {
	return nil
}

// Scan mocked rows
func (m *rowsMock) Scan(dest ...any) error {
	mr := m.Rows[m.index]

	if len(mr) != len(dest) {
		return fmt.Errorf("invalid len of dest")
	}

	for index, value := range mr {
		v := reflect.ValueOf(dest[index])
		v.Elem().Set(reflect.ValueOf(value))
	}

	m.index++
	return nil
}
