package part2

import (
	"testing"
)

func TestResult(t *testing.T) {
	result1 := Result("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	result2 := Result("bvwbjplbgvbhsrlpgdmjqwftvncz")
	result3 := Result("nppdvjthqldpwncqszvftbrmjlhg")
	result4 := Result("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	result5 := Result("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")

	if result1 != 19 {
		t.Errorf("Expected 19, got %d", result1)
	}

	if result2 != 23 {
		t.Errorf("Expected 23, got %d", result2)
	}

	if result3 != 23 {
		t.Errorf("Expected 23, got %d", result3)
	}

	if result4 != 29 {
		t.Errorf("Expected 29, got %d", result4)
	}

	if result5 != 26 {
		t.Errorf("Expected 26, got %d", result5)
	}

}
