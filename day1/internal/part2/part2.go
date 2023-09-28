package part2

import (
	"sort"
	"strconv"
	"strings"
)

func Result(input string) int {
    elves := strings.Split(input, "\n\n")
    totals := []int{}
    result := 0


    for _, elf := range(elves) {
        total := calculateElf(elf)
        totals = append(totals, total)
    }

    sort.Slice(totals, func(i, j int) bool {
        if (totals[i] > totals[j]) {
            return true
        }
        return false
    })

    for _, total := range(totals[:3]) {
        result += total
    }

    return result
}

func calculateElf(elf string) int {
    total := 0
    numbers := strings.Split(elf, "\n")

    for _, number := range(numbers) {
        n, _ := strconv.Atoi(number)
        total += n
    }

    return total
}

