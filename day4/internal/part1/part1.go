package part1

import (
	"strconv"
	"strings"
)

func Result(input string) int {

	result := 0
	assignments := strings.Split(input, "\n")

	for _, assignment := range assignments {
		if assignment != "" {
			result += calculate(assignment)
		}
	}

	return result
}

func calculate(assignment string) int {
	group := strings.Split(assignment, ",")
	values := []int{}

	for _, x := range group {
		v := strings.Split(x, "-")
		for _, y := range v {
			a, _ := strconv.Atoi(y)
			values = append(values, a)
		}
	}

	if (values[0] <= values[2] && values[1] >= values[3]) || (values[2] <= values[0] && values[3] >= values[1]) {
		return 1
	}

	return 0
}
