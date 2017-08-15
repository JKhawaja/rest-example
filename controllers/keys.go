package controllers

import (
	"fmt"

	"github.com/JKhawaja/replicated/app"
	"github.com/JKhawaja/replicated/services/github"
	. "github.com/JKhawaja/replicated/util"

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
	var response []User

	// check that username has been provided (unecessary as already checked by goa)
	if len(ctx.Payload) < 1 {
		return ctx.BadRequest(fmt.Errorf("Please provide a username."))
	}

	// remove any duplicate names in request
	names := RemoveDuplicates(ctx.Payload)

	// get keys for each username
	for _, name := range names {
		keys, err := c.Client.ListKeys(name)
		if err != nil {
			goalogrus.Entry(ctx).Errorf("GitHub API Access error: %+v", err)
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
