package part2

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

	if result != 4 {
		t.Errorf("Expected 4, got %d\n", result)
	}
}

