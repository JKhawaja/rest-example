package controllers

import (
	"fmt"

	"github.com/JKhawaja/rest-example/controllers/app"
	"github.com/JKhawaja/rest-example/services/github"
	. "github.com/JKhawaja/rest-example/util"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/logrus"
)

// KeysController implements the keys resource.
type KeysController struct {
	*goa.Controller
	Client github.Client
}

// User is the type for GitHub user's...
type User struct {
	Name string         `json:"username"`
	Keys []*app.UserKey `json:"keys"`
}

// NewKeysController creates a keys controller.
func NewKeysController(service *goa.Service, client github.Client) *KeysController {
	return &KeysController{
		Controller: service.NewController("KeysController"),
		Client:     client,
	}
}

// List runs the list action.
func (c *KeysController) List(ctx *app.ListKeysContext) error {

	// check if GitHub is up
	if !c.Client.GetStatus() {
		err := fmt.Errorf("GitHub may be temporarily down. Please try again.")
		return ctx.GatewayTimeout(err)
	}

	var response []User

	// check that username has been provided
	if len(ctx.Payload) < 1 {
		return ctx.BadRequest(fmt.Errorf("Please provide a username."))
	}

	// remove any duplicate names in request
	names := RemoveDuplicates(ctx.Payload)

	// get keys for each username
	for _, name := range names {
		keys, err := c.Client.ListKeys(name)
		if err != nil {
			if _, ok := ctx.Context.Value("test").(bool); !ok {
				goalogrus.Entry(ctx).Errorf("GitHub API access error: %+v", err)
			}
			return ctx.InternalServerError()
		}

		newKeys := ConvertList(keys)

		u := User{
			Name: name,
			Keys: newKeys,
		}

		response = append(response, u)
	}

	// create and write response
	res := app.UserCollection{}
	for _, un := range response {
		newUser := &app.User{
			Username: un.Name,
			Keys:     un.Keys,
		}
		res = append(res, newUser)
	}

	return ctx.OK(res)
}
