// +build gofuzz

package fuzz

import (
	"github.com/JKhawaja/rest-example/util"
)

func FuzzRemoveDuplicates(data []byte) int {
	slice := []string{string(data), string(data)}
	newSlice := util.RemoveDuplicates(slice)
	if len(newSlice) != 1 {
		return 1
	}

	return 0
}

func FuzzNameVerification(data []byte) int {
	slice := []string{string(data)}
	_, ok := util.NameVerification(slice)
	if !ok {
		return 1
	}

	return 0
}
