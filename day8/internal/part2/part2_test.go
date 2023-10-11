package part2

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

	if result != 8 {
		t.Errorf("Expected 8, got %d\n", result)
	}
}
