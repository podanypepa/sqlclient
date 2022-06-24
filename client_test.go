package sqlclient_test

import (
	"fmt"
	"testing"

	"github.com/podanypepa/sqlclient"
	"github.com/stretchr/testify/assert"
)

func TestMocking(t *testing.T) {
	t.Run("example", func(t *testing.T) {

		sqlclient.StartMockServer()

		client, err := sqlclient.Open("mysql", "jaba daba duuu")
		assert.NoError(t, err, "Open err")

		sqlclient.AddMock(sqlclient.Mock{
			Query: "SELECT id, name FROM users WHERE id=?;",
			Args:  []any{100},
			Err:   nil,
			Cols:  []string{"id, name"},
			Rows: [][]any{
				{100, "pepa"},
				{100, "josef"},
			},
		})

		q := "SELECT id, name FROM users WHERE id=?;"
		res, err := client.Query(q)

		if !assert.NoErrorf(t, err, "Open err: %v", err) {
			t.FailNow()
		}

		assert.NotNil(t, res, "res is nil")

		type data struct {
			ID   int
			Name string
		}
		for res.Next() {
			var d data
			err := res.Scan(&d.ID, &d.Name)
			assert.NoError(t, err, "Scan err")
			fmt.Printf("TEST id: %d, name: %s\n", d.ID, d.Name)
			assert.Equal(t, 100, d.ID, "must be 100")
		}
	})
}
