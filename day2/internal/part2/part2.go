package part2

import "strings"

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
const (
	opponent = 0
	me       = 1
)

func Result(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // remove empty line at the end
	total := 0
	for _, line := range lines {
		total += evaluate(line)
	}

	return total
}

func evaluate(line string) int {

	rules := [3]int{lose, draw, victory}
	points := map[int]int{
		lose:    0,
		draw:    3,
		victory: 6,
	}

	choicePoints := map[int]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}
	options := strings.Fields(line)

	a := int(options[me][0]) - int('X')
	b := int(options[opponent][0]) - int('A')
	result := rules[a]
    choice := choose(b, result)

	return points[result] + choicePoints[choice]
}

func choose(opponentChoice int, result int) int {
    if(result == lose && opponentChoice == rock) {
        return scissors
    } else if( result == victory && opponentChoice == scissors) {
        return rock
    }

    return opponentChoice + result
}
