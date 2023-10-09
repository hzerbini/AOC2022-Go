package part1

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	From     int
	To       int
	Quantity int
}

func Result(input string) string {
	crates, instructions := parse(input)
	result := ""

	for _, instruction := range instructions {
		for i := 0; i < instruction.Quantity; i++ {
			row := crates[instruction.From]
			crates[instruction.From] = row[:len(row)-1]
			crates[instruction.To] = append(crates[instruction.To], row[len(row)-1])
		}
	}

	for _, crate := range crates {
		if len(crate) > 0 {
			result += string(crate[len(crate)-1])
		}
	}

	return result
}

func parse(input string) ([][]rune, []Instruction) {
	entries := strings.SplitN(input, "\n\n", 2)

	if len(entries) != 2 {
		panic("Wrong number of entries to parse function")
	}

	instructions := parseInstructions(entries[1])
	crates := parseCrates(entries[0])

	return crates, instructions
}

func parseInstructions(input string) []Instruction {
	instructions := []Instruction{}

	for i, line := range strings.Split(input, "\n") {
		if line != "" {
			line = strings.ReplaceAll(line, "move ", "")
			line = strings.ReplaceAll(line, "from ", "")
			line = strings.ReplaceAll(line, "to ", "")
			arr := []int{}
			for y, n := range strings.Split(line, " ") {
				if y >= 3 {
					panic(fmt.Errorf(
						"Instruction (%d) with wrong parameters count",
						i,
					))
				}

				num, err := strconv.Atoi(n)
				if err != nil {
					panic(fmt.Errorf(
						"Error converting string('%s') to number",
						n,
					))
				}

				arr = append(arr, num)
			}

			instruction := Instruction{
				arr[1] - 1,
				arr[2] - 1,
				arr[0],
			}

			instructions = append(instructions, instruction)
		}
	}

	return instructions
}

func parseCrates(input string) [][]rune {
	crates := [][]rune{}
	lines := strings.Split(input, "\n")

	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		lineCrates := splitCrates(line)
		if len(lines)-1 == i {
			for j := 0; j < len(lineCrates); j++ {
				crates = append(crates, []rune{})
			}

		} else {
			for j, crate := range lineCrates {
				if crate != ' ' {
					crates[j] = append(crates[j], crate)
				}
			}
		}
	}

	return crates
}

func splitCrates(input string) []rune {
	crates := []rune{}

	for i := 0; i < len(input); i++ {
		if i%4 == 3 {
			crates = append(crates, trimCrate(input[i-3:i]))
		} else if i == len(input)-1 {
			crates = append(crates, trimCrate(input[i-i%4:i]))

		}
	}

	return crates
}

func trimCrate(crate string) rune {
	crate = strings.Trim(crate, " ")
	crate = strings.Trim(crate, "[")
	crate = strings.Trim(crate, "]")

	if len(crate) == 0 {
		return ' '
	}
	return rune(crate[0])
}
