package part2

import (
	"strconv"
	"strings"
)

func Result(input string) int {
	trees := parse(input)
	result := 0

	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			score := calculateScenicScore(trees, x, y)
			if score > result {
				result = score
			}
		}
	}

	return result
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

func calculateScenicScore(list [][]int, x int, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	// count left
	left := 0
	for i := x - 1; i >= 0; i-- {
		left++
		if list[y][i] >= list[y][x] {
			break
		}
	}

	// count right
	right := 0
	for i := x + 1; i < len(list[y]); i++ {
		right++
		if list[y][i] >= list[y][x] {
			break
		}
	}

	// count top
	top := 0
	for i := y - 1; i >= 0; i-- {
		top++
		if list[i][x] >= list[y][x] {
			break
		}
	}

	// count bottom
	bottom := 0
	for i := y + 1; i < len(list); i++ {
		bottom++
		if list[i][x] >= list[y][x] {
			break
		}
	}

	return left * right * top * bottom
}
