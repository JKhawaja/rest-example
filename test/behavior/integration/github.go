// +build integration

package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/JKhawaja/rest-example/client"
	"github.com/JKhawaja/rest-example/services/github"
	"github.com/JKhawaja/rest-example/test"

	. "github.com/smartystreets/goconvey/convey"
)

// Integration Test
func TestClient(t *testing.T) {

	// start NewMockServer()
	mockServer := test.NewMockServer()
	log.Fatal(mockServer.ListenAndServe())

	// create Github mock Client
	mockClient := &github.MockClient{}

	// create service client
	serverClient := client.New(nil)

	// Test Table
	path := "localhost:9090" + client.ListKeysPath()
	payloads := [][]string{
		{"tom", "dave"},
		{"john", "john", "john"},
		{"1234", "5678", "9101112"},
	}
	testData := []test.Case{
		{Context: context.Background(), Path: path, Payload: payloads[0]},
		{Context: context.Background(), Path: path, Payload: payloads[1]},
		{Context: context.Background(), Path: path, Payload: payloads[2]},
	}

	// Program mock client responses (TODO: according to test table)
	mockClient.On("ListKeys", "david").Return([]github.Key{{ID: 1, Key: "david"}}, nil)
	mockClient.On("ListKeys", "tom").Return([]github.Key{{ID: 2, Key: "tom"}}, nil)
	mockClient.On("ListKeys", "1234").Return([]github.Key{}, fmt.Errorf("Incorrect Username"))

	// Test Case #1
	Convey("Given a HTTP request for /keys", t, func() {
		req, err := serverClient.NewListKeysRequest(testData[0].Context, testData[0].Path, testData[0].Payload)
		if err != nil {
			t.Fatalf("GitHub Client - ListKeys method - Test Case #1 failed with: %+v", err)
		}

		Convey("When the request is sent by the client", func() {

			resp, err := serverClient.Client.Do(testData[0].Context, req)
			if err != nil {
				t.Fatalf("GitHub Client - ListKeys method - Test Case #1 failed with: %+v", err)
			}

			mockServer.Stop(1 * time.Second)

			Convey("Then the response should be a 200", func() {
				So(resp.StatusCode, ShouldEqual, 200)

				var response github.Keys
				if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
					t.Fatalf("GitHub Client - ListKeys method - Test Case #1 failed with: %+v", err)
				}
				So(response[0].ID, ShouldEqual, 1)
				So(response[0].Key, ShouldEqual, "david")
			})
		})
	})
}
