package part2_test

import (
	"hzerbini/day1/internal/part2"
	"testing"
    "os"
)

func TestResult(t *testing.T) {
    input, _ := os.ReadFile("../../input.test")

    result := part2.Result(string(input))

    if(result != 45000) {
        t.Errorf("Expects 45000, got %d", result)
    }

}
