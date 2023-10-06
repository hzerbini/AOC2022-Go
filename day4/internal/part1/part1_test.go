package part1

import (
	"os"
	"testing"
)

func TestResult(t *testing.T) {
	input, err := os.ReadFile("../../input.dev")
	if err != nil {
		panic(err)
	}

	result := Result(string(input))

	if result != 2 {
		t.Errorf("Expected 2, got %d\n", result)
	}
}
