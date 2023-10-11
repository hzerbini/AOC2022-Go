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

	if result != 21 {
		t.Errorf("Expected 21, got %d\n", result)
	}
}
