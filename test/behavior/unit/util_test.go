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
