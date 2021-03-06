package action

import (
	"testing"
)

func TestSearchSuccess(t *testing.T) {
	result, _ := Search("match")
	correctValue := (interface{})("No match.")
	if result[0] != correctValue {
		t.Fatal("failed test")
	}
}

func TestSearchSuccess2(t *testing.T) {
	result, _ := Search("match2")
	correctValue := (interface{})("No match.")
	if result[0] != correctValue {
		t.Fatal("failed test")
	}
}

func TestSearchFailed(t *testing.T) {
	result, _ := Search("match3")
	incorrectValue := (interface{})("match!")
	if result[0] == incorrectValue {
		t.Fatal("failed test")
	}
}
