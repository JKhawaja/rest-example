package util

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/JKhawaja/rest-example/controllers/app"
	"github.com/JKhawaja/rest-example/services/github"
)

var usernameRegex *regexp.Regexp

func init() {
	// FIXME: regex does not detect length (failing test)
	usernameRegex = regexp.MustCompile(`^[a-z](?:[a-z0-9]|-(?:[a-z0-9])).{0,38}$`)
}

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

// NameVerification ...
func NameVerification(names []string) (string, bool) {
	for _, s := range names {
		if !usernameRegex.MatchString(s) {
			return s, false
		}
	}

	return "", true
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
