// +build integration

package github

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JKhawaja/rest-example/client"
	"github.com/JKhawaja/rest-example/test"

	. "github.com/smartystreets/goconvey/convey"
)

// Integration Test
func TestClient(t *testing.T) {

	// start NewMockServer()
	mockServer := test.NewMockServer()

	// create Github mock Client
	mockClient := &MockClient{}

	// create service client
	serverClient := client.New(nil)

	// Test Table
	path := "localhost:8080" + client.ListKeysPath()
	payloads := [][]string{
		{"tom", "dave"},
		{"john", "john", "john"},
		{"1234", "5678", "9101112"},
	}
	testData := []test.Case{
		{Context: context.Background(), Path: path, Payload: payloads[0]},
		{context.Background(), path, payloads[1]},
		{context.Background(), path, payloads[2]},
	}

	// Program mock client responses (TODO: according to test table)
	mockClient.On("ListKeys", "david").Return([]Key{Id: 1, Key: "david"}, nil)
	mockClient.On("ListKeys", "tom").Return([]Key{Id: 2, Key: "tom"}, nil)
	mockClient.On("ListKeys", "456").Return([]Key{}, fmt.Errorf("Incorrect Username"))

	// Test Case #1
	Convey("Given a HTTP request for /keys", t, func() {
		req, err := serverClient.NewListKeysRequest(testData[0], path, payload)
		if err != nil {
			t.Fatalf("GitHub Client - ListKeys method - Test Case #1 failed with: %+v", err)
		}

		Convey("When the request is sent by the client", func() {

			resp, err := serverClient.Client.Do(ctx, req)
			if err != nil {
				t.Fatalf("GitHub Client - ListKeys method - Test Case #1 failed with: %+v", err)
			}

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)

				var response Keys
				json.Unmarshal(resp.Body.Bytes(), &response)
				So(response[0].ID, ShouldEqual, 1)
				So(response[0].Key, ShouldEqual, "david")
			})
		})
	})
}
