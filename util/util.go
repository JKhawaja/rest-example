package util

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/JKhawaja/rest-example/controllers/app"
	"github.com/JKhawaja/rest-example/services/github"
)

// RemoveDuplicates ...
func RemoveDuplicates(names []string) []string {
	seen := map[string]bool{}
	result := []string{}

	for n := range names {
		if seen[names[n]] == true {
			//do nothing
		} else {
			seen[names[n]] = true
			result = append(result, names[n])
		}
	}

	return result
}

// ConvertList ...
func ConvertList(list []github.Key) []*app.UserKey {
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

// BuildBinary ...
// FIXME: does not work on Windows
func BuildBinary(name string) (string, error) {
	// directory
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}

	// command
	binName := filepath.Join(tmpDir, name)
	staticLink := `'-extldflags "-static"'`
	command := []string{
		"go", "build", "-o", binName, "-a", "--ldflags", staticLink, ".",
	}
	cmd := exec.Command(command[0], command[1:]...)

	// envVar
	gopath := os.Getenv("GOPATH")
	cmd.Env = []string{
		"GOOS=linux",
		"GOARCH=amd64",
		"GOPATH=" + gopath,
	}

	// execute
	data, err := cmd.CombinedOutput()
	if err != nil {
		return string(data), err
	}

	return binName, nil
}
