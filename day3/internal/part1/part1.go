package part1

import (
	"slices"
	"strings"
)

func Result(input string) int {
    total := 0
	rucksacks := strings.Split(input, "\n")
	for _, rucksack := range rucksacks {
		total += priority([]byte(rucksack))
	}

	return total
}

func priority(rucksack []byte) int {
	length := len(rucksack)
	middle := length / 2
	partition1 := rucksack[:middle]
	partition2 := rucksack[middle:]
	for _, c := range partition1 {
		if slices.Contains(partition2, c) {
            return calculate(c)
		}
	}

	return 0
}

func calculate(item byte) int {
	if item >= 97 && item <= 122 {
        return int(item - 96)
	} else if item >= 65 && item <= 90 {
        return int(item - 38)
    }
    return 0
}
