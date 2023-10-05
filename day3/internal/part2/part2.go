package part2

import (
	"fmt"
	"strings"
)

func Result(input string) int {
	total := 0
	rucksacks := strings.Split(input, "\n")
	group := []string{}
	for i, rucksack := range rucksacks {
		group = append(group, rucksack)

		if (i+1)%3 == 0 {
			p, err := priority(group)
			if err != nil {
				panic(err)
			}
			total += p
			group = []string{}
		}
	}

	return total
}

func priority(group []string) (int, error) {
	length := len(group)
	if length != 3 {
		return 0, fmt.Errorf("Wrong quantity of rucksacks per group. Expected 3, got %d", length)
	}
	rucksack1 := group[0]
	rucksack2 := group[1]
	rucksack3 := group[2]

	for _, item := range rucksack1 {
		if strings.ContainsRune(rucksack2, item) && strings.ContainsRune(rucksack3, item) {
			return calculate(byte(item)), nil
		}
	}

	return 0, nil
}

func calculate(item byte) int {
	if item >= 97 && item <= 122 {
		return int(item - 96)
	} else if item >= 65 && item <= 90 {
		return int(item - 38)
	}
	return 0
}
