package part1

import (
	"os"
	"testing"
)

func TestResult(t *testing.T) {

    input, err := os.ReadFile("../../input.test")

    if(err != nil) {
        t.Error(err)
    }

    result := Result(string(input))

    if(result != 15) {
        t.Errorf("Expected 15, got %d", result)
    }
}

func TestEvaluate(t *testing.T) {
    test1 := evaluate("A Y")
    test2 := evaluate("B X")
    test3 := evaluate("C Z")

    if(test1 != 8) {
        t.Errorf("test1: Expected 8, got %d", test1)
    }

    if(test2 != 1) {
        t.Errorf("test2: Expected 1, got %d", test2)
    }

    if(test3 != 6) {
        t.Errorf("test3: Expected 6, got %d", test3)
    }
}
