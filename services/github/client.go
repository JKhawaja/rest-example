package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GHC ...
type GHC struct {
	Client *http.Client
}

// NewGHC ...
func NewGHC() Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 10 * time.Second,
	}

	return &GHC{
		Client: &http.Client{Transport: tr},
	}
}

// ListKeys ...
func (g *GHC) ListKeys(username string) ([]Key, error) {
	emptyResp := []Key{}

	url := fmt.Sprintf("http://api.github.com/users/%s/keys", username)
	resp, err := g.Client.Get(url)
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
