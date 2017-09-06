// +build unit

package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/JKhawaja/rest-example/client"
)

// TestListKeys is a basic unit test example using the generated goa client code
// NOTE: can also make use of the generated `test` package (located inside the app directory)
func TestListKeys(t *testing.T) {

	tests := [][]string{
		{"tom", "dave"},
		{"john", "john", "john"},
		{"1234", "5678", "9101112"},
	}

	for _, usernames := range tests {
		keysClient := client.New(nil)

		ctx := context.Background()
		path := "localhost:8080" + client.ListKeysPath()
		resp, err := keysClient.ListKeys(ctx, path, usernames)
		if err != nil {
			t.Fatalf("Query %s failed with: %+v", usernames, err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Failed request: %s", usernames)
		}

		resp.Body.Close()
	}
}
