package part1

import "strings"

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

	rules := [3][3]int{{draw, lose, victory}, {victory, draw, lose}, {lose, victory, draw}} // 0 -> rock, 1 -> paper, 2 -> scissors
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
	result := rules[a][b]

	return points[result] + choicePoints[a]
}
