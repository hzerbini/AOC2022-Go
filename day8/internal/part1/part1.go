package part1

import (
	"strconv"
	"strings"
)

func Result(input string) int {
	trees := parse(input)
	count := 0

	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			if isVisible(trees, x, y) {
				count++
			}
		}
	}

	return count
}

func parse(input string) [][]int {
	result := [][]int{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line != "" {
			aux := []int{}

			for _, c := range line {
				n, err := strconv.Atoi(string(c))
				if err != nil {
					panic("Tree high is not a number")
				}

				aux = append(aux, n)
			}

			result = append(result, aux)
		}
	}

	return result
}

func isVisible(list [][]int, x int, y int) bool {
	// Visible at the edge
	if x == 0 || y == 0 || x == len(list[y])-1 || y == len(list)-1 {
		return true
	}

	// Is visible from the left
	max := 0
	for i := 0; i < x; i++ {
		if max < list[y][i] {
			max = list[y][i]
		}
	}
	if max < list[y][x] {
		return true
	}

	// Is visible from the right
	max = 0
	for i := x + 1; i < len(list[y]); i++ {
		if max < list[y][i] {
			max = list[y][i]
		}
	}
	if max < list[y][x] {
		return true
	}

	// Is visible from the top
	max = 0
	for i := 0; i < y; i++ {
		if max < list[i][x] {
			max = list[i][x]
		}
	}
	if max < list[y][x] {
		return true
	}

	// Is visible from the bottom
	max = 0
	for i := y + 1; i < len(list); i++ {
		if max < list[i][x] {
			max = list[i][x]
		}
	}
	if max < list[y][x] {
		return true
	}

	return false
}
