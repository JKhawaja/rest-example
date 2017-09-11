// +build unit

package unit

import (
	"testing"

	"github.com/JKhawaja/rest-example/util"
)

func TestRemoveDuplicates(t *testing.T) {

	duplicates := [][]string{
		{"1", "1", "1"},
		{"1", "1", "2", "2"},
		{"1", "2", "2", "3", "3", "3"},
	}

	for i, d := range duplicates {
		ret := util.RemoveDuplicates(d)

		if len(ret) > i+1 {
			t.Fatal("Util Test - RemoveDuplicates - failed to remove all duplicates.")
		}
	}
}

func TestNameVerification(t *testing.T) {
	invalidNames := [][]string{
		{"1wrong", "2wrong", "3wrong"},
		{"correct", "still-correct", "still-correct1234"},
		{"incorrect", "still--incorrect", "-very--incorrect-"},
		{"this-one-is-just-way-too-long-to-be-a-valid-github-username-1234"},
	}

	for i, names := range invalidNames {
		_, ret := util.NameVerification(names)

		if i != 1 && ret {
			t.Fatalf("Util Test - NameVerification - failed to mark incorrect username as incorrect: %+v", names)
		}
	}

}
