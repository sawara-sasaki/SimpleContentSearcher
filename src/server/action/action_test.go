package action

import (
	"testing"
)

func TestSearchSuccess(t *testing.T) {
	result, _ := Search("")
	correctValue := (interface{})("No match.")
	if result[0] != correctValue {
		t.Fatal("failed test")
	}
}
