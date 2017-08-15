package util

import (
	"github.com/JKhawaja/replicated/app"
	"github.com/JKhawaja/replicated/services/github"
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
