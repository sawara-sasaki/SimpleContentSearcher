package action

import (
	"testing"
)

func TestSearchSuccess(t *testing.T) {
	result, _ := Search("http://match")
	correctValue := (interface{})("No match.")
	if result[0] != correctValue {
		t.Fatal("failed test")
	}
}

func TestSearchSuccess2(t *testing.T) {
	result, _ := Search("http://match2")
	correctValue := (interface{})("No match.")
	if result[0] != correctValue {
		t.Fatal("failed test")
	}
}

func TestSearchFailed(t *testing.T) {
	result, _ := Search("https://www.youtube.com/c/google/videos")
	incorrectValue := (interface{})("match!")
	if len(result) > 0 && result[0] == incorrectValue {
		t.Fatal("failed test")
	}
}

func TestSearchFailed2(t *testing.T) {
	result, _ := Search("https://www.youtube.com/watch?v=ZRCdORJiUgU")
	incorrectValue := (interface{})("match!")
	if len(result) > 0 && result[0] == incorrectValue {
		t.Fatal("failed test")
	}
}

func TestSearchFailed3(t *testing.T) {
	result, _ := Search("https://www.youtube.com/results?search_query=google")
	incorrectValue := (interface{})("match!")
	if len(result) > 0 && result[0] == incorrectValue {
		t.Fatal("failed test")
	}
}
