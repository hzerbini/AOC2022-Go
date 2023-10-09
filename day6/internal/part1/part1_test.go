package part1

import (
	"os"
	"testing"
)

func TestResult(t *testing.T) {
	input1, err := os.ReadFile("../../input.test")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	input2, err := os.ReadFile("../../input2.test")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	input3, err := os.ReadFile("../../input3.test")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	input4, err := os.ReadFile("../../input4.test")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	result1 := Result(string(input1))
	result2 := Result(string(input2))
	result3 := Result(string(input3))
	result4 := Result(string(input4))

	if result1 != 5 {
		t.Errorf("Expected 5, got %d", Result(string(input1)))
	}

	if result2 != 6 {
		t.Errorf("Expected 6, got %d", Result(string(input2)))
	}

	if result3 != 10 {
		t.Errorf("Expected 10, got %d", Result(string(input3)))
	}

	if result4 != 11 {
		t.Errorf("Expected 11, got %d", Result(string(input4)))
	}

}
