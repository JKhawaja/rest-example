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
