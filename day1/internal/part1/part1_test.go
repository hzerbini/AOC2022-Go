package part1_test

import (
	"testing"
	"hzerbini/day1/internal/part1"
    "os"
)

func TestPart1(t *testing.T) {
    inputBytes, _ := os.ReadFile("../../input.test")
    input := string(inputBytes)

    result := part1.Result(input)
    if(result != 24000) {
        t.Errorf("Expects 24000, got %d", result)
    }
}
