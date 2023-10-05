package part1

import (
	"os"
	"testing"
)

func TestResult(t *testing.T) {
	input, err := os.ReadFile("../../input.test")
	if err != nil {
		panic("Error reading test input file")
	}

	result := Result(string(input))
	if result != 157 {
		t.Errorf("Expected 157, got %d", result)
	}
}
