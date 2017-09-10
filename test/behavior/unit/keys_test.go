// +build unit

package unit

import (
	"context"
	"fmt"
	"testing"

	"github.com/JKhawaja/rest-example/controllers"
	"github.com/JKhawaja/rest-example/controllers/app/test"
	"github.com/JKhawaja/rest-example/services"
	"github.com/JKhawaja/rest-example/services/github"
	"github.com/JKhawaja/rest-example/test/mock"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/logrus"
	"github.com/sirupsen/logrus"
)

// TestListKeys ...
func TestListKeys(t *testing.T) {

	// Real Service
	service1 := goa.New("Test")
	status := services.NewStatus()
	ghc := github.NewClient(status)
	ctrlr1 := controllers.NewKeysController(service1, ghc)

	// Mock Service
	service2 := goa.New("Test")
	logger := logrus.New()
	service2.WithLogger(goalogrus.New(logger))
	mockClient := &mock.GithubClient{}
	mockClient.On("ListKeys", "david").Return([]github.Key{{ID: 1, Key: "david"}}, nil)
	mockClient.On("ListKeys", "tom").Return([]github.Key{{ID: 2, Key: "tom"}}, nil)
	mockClient.On("ListKeys", "error").Return([]github.Key{{ID: 3, Key: "error"}}, fmt.Errorf("This is an error!"))
	mockGHC := github.Client(mockClient)
	ctrlr2 := controllers.NewKeysController(service2, mockGHC)

	payloads := [][]string{
		{"david", "tom"},          // 200
		{},                        // 400
		{"tom"},                   // 500
		{"tom", "david", "error"}, // 504
	}

	// OK Test
	_, coll := test.ListKeysOK(t, nil, service2, ctrlr2, payloads[0])
	if coll[0].Keys[0].Key != "david" || coll[1].Keys[0].Key != "tom" {
		t.Fatalf("ListKeysOK test failed with wrong collection value.")
	}

	// Bad Request Test
	_, err := test.ListKeysBadRequest(t, nil, service2, ctrlr2, payloads[1])
	if err.Error() != "Please provide a username." {
		t.Fatalf("ListKeysBadRequest test errored out with: %+v ", err)
	}

	// Gateway Timeout Error
	err = status.Set("github", false)
	if err != nil {
		t.Fatalf("ListKeysGatewayTimout error: %+v", err)
	}
	_, err = test.ListKeysGatewayTimeout(t, nil, service1, ctrlr1, payloads[2])
	if err.Error() != "GitHub may be temporarily down. Please try again." {
		t.Fatalf("ListKeysGatewayTimeout test errored out with: %+v ", err)
	}

	// Internal Server Error
	parent := context.Background()
	ctx := context.WithValue(parent, "test", true)
	_ = test.ListKeysInternalServerError(t, ctx, nil, ctrlr2, payloads[3])

}
