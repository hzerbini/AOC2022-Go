package part2

import (
	"os"
	"testing"
)

func TestResult(t *testing.T) {
	input, err := os.ReadFile("../../input.test")

	if err != nil {
		t.Error(err)
	}

	result := Result(string(input))

	if result != 12 {
		t.Errorf("Expected 12, got %d", result)
	}
}

func TestEvaluate(t *testing.T) {
	test1 := evaluate("A Y")
	test2 := evaluate("B X")
	test3 := evaluate("C Z")

	if test1 != 4 {
		t.Errorf("test1: Expected 4, got %d", test1)
	}

	if test2 != 1 {
		t.Errorf("test2: Expected 1, got %d", test2)
	}

	if test3 != 7 {
		t.Errorf("test3: Expected 7, got %d", test3)
	}
}

func TestChoose(t *testing.T) {
	const (
		lose    = -1
		draw    = 0
		victory = 1
	)
	const (
		rock     = 0
		paper    = 1
		scissors = 2
	)

	test1 := choose(rock, draw)
	test2 := choose(paper, lose)
	test3 := choose(scissors, victory)

	if test1 != rock {
        t.Errorf("test1: Expected rock, got %d", test1)
	}

	if test2 != rock {
        t.Errorf("test2: Expected rock, got %d", test2)
	}

	if test3 != rock {
        t.Errorf("test3: Expected rock, got %d", test3)
	}
}
