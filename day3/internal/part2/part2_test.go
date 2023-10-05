package part2

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
	if result != 70 {
		t.Errorf("Expected 70, got %d", result)
	}
}
