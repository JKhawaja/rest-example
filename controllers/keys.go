package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JKhawaja/replicated/app"

	"github.com/goadesign/goa"
)

// KeysController implements the keys resource.
type KeysController struct {
	*goa.Controller
	Client *http.Client
}

// Key ..
type Key struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// User is the type for GitHub user's...
type User struct {
	Name string         `json:"username"`
	Keys []*app.UserKey `json:"keys"`
}

// NewKeysController creates a keys controller.
func NewKeysController(service *goa.Service, client *http.Client) *KeysController {
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
	names := removeDuplicates(ctx.Payload)

	// get keys for each username
	for _, name := range names {
		keys, err := getGitHubKeys(name, c.Client)
		if err != nil {
			return ctx.BadRequest(err)
		}

		newKeys := convertList(keys)

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

func removeDuplicates(names []string) []string {
	seen := map[string]bool{}
	result := []string{}

	for n := range names {
		if seen[names[n]] == true {
			//do nothing
		} else {
			// record element as seen
			seen[names[n]] = true
			// append to new slice
			result = append(result, names[n])
		}
	}
	// return unique slice
	return result
}

// this will return a list of Keys for a username unless an error occurs
func getGitHubKeys(username string, client *http.Client) ([]Key, error) {
	emptyResp := []Key{}

	url := fmt.Sprintf("http://api.github.com/users/%s/keys", username)
	resp, err := client.Get(url)
	if err != nil {
		return emptyResp, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Error reading GitHub response body: %+v for username: %s", err, username)
		return emptyResp, err
	}

	var response []Key
	err = json.Unmarshal(b, &response)
	if err != nil {
		err := fmt.Errorf("Could not decode response format: %+v, for username: %s", err, username)
		return emptyResp, err
	}

	return response, nil
}

func convertList(list []Key) []*app.UserKey {
	var newList []*app.UserKey

	for _, k := range list {
		uk := &app.UserKey{
			ID:  k.ID,
			Key: k.Key,
		}

		newList = append(newList, uk)
	}

	return newList
}
