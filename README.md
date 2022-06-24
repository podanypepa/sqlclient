# sqlclient
Go SQL client

## Why?

SQL client for easy mocking connection to the SQL server and unit testing.

## 1. Testing example

You can write standart tests with started Mock sever which separate you from
fyzical server.

```go
// client_test.go

import (
	"fmt"
	"testing"
	
	"github.com/podanypepa/sqlclient"
	"github.com/stretchr/testify/assert"
)

func TestMocking(t *testing.T) {
	t.Run("example", func(t *testing.T) {

		// start Mock server
		sqlclient.StartMockServer()

		// open sql connection like always
		client, err := sqlclient.Open("mysql", "jaba daba duuu")
		assert.NoError(t, err, "Open err")

		// mocking sql query with params and response
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

		// make real query to the server
		q := "SELECT id, name FROM users WHERE id=?;"
		res, err := client.Query(q)

		// finaly asseritng result
		if !assert.NoErrorf(t, err, "Open err: %v", err) {
			t.FailNow()
		}

		assert.NotNil(t, res, "res is nil")

		// you can iterate over mocked data... :)
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
```

## 2. Mocking directly in app

You are able start mock server and mock data directly in the your app.
Then you can run app with real db queryies without real connection to the server.

Suma sumarum: you can add to the app anywhere where you need mocked data and 
run real queries without changes in logic of app.

```go
// main.go

import (
	"fmt"
	"log"

	"github.com/podanypepa/sqlclient"
)

func init() {
	// comment next line if you want run against real sql server
	sqlclient.StartMockServer()
}

func main() {

	c := sqlclient.Open("mysql", "abraca dabra :)")

	// comment next line if you want run against real sql server
	addMocks()

	q := "SELECT id, name FROM users WHERE id=?;"
	res, err := client.Query(q)
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var d data
		err := res.Scan(&d.ID, &d.Name)
		fmt.Printf("TEST id: %d, name: %s\n", d.ID, d.Name)
	}		
}

func addMocks() {
	// add mocked query
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
}

```
