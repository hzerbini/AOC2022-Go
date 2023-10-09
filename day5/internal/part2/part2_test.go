package part2

import (
	"os"
	"testing"
)

func TestResult(t *testing.T) {
	input, err := os.ReadFile("../../input.test")
	if err != nil {
		panic(err)
	}

	result := Result(string(input))

	if result != "MCD" {
		t.Errorf("Expected \"CMZ\", got %s\n", result)
	}
}
