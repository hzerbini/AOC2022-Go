package part1

import (
	"os"
	"testing"
)

func TestResult(t *testing.T) {
	input, err := os.ReadFile("../../input.test")
	if err != nil {
		t.Fatal(err)
	}
	result := Result(string(input))

	if result != 95437 {
		t.Errorf("Expected 95437, got %d", result)
	}
}
