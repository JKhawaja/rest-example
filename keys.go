package main

import (
	"github.com/JKhawaja/replicated/app"
	"github.com/goadesign/goa"
)

// KeysController implements the keys resource.
type KeysController struct {
	*goa.Controller
}

// NewKeysController creates a keys controller.
func NewKeysController(service *goa.Service) *KeysController {
	return &KeysController{Controller: service.NewController("KeysController")}
}

// List runs the list action.
func (c *KeysController) List(ctx *app.ListKeysContext) error {
	// KeysController_List: start_implement

	// Put your logic here

	// KeysController_List: end_implement
	res := app.KeysCollection{}
	return ctx.OK(res)
}

// Key ..
type Key struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// this will return a list of Keys for a username unless an error occurs
func getGitHubKeys(username string) ([]Key, error) {

}
