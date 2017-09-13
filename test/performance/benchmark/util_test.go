// +build benchmark

package benchmark

import (
	"testing"

	"github.com/JKhawaja/rest-example/util"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
	names := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for n := 0; n < b.N; n++ {
		util.RemoveDuplicates(names)
	}
}

func BenchmarkNameVerification(b *testing.B) {
	invalidNames := [][]string{
		{"1wrong", "2wrong", "3wrong"},
		{"incorrect", "still--incorrect", "-very--incorrect-"},
		{"this-one-is-just-way-too-long-to-be-a-valid-github-username-1234"},
	}

	for n := 0; n < b.N; n++ {
		for _, names := range invalidNames {
			util.NameVerification(names)
		}
	}
}
