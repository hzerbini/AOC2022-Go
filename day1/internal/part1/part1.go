package part1

import (
    "strings"
    "strconv"
)

func Result(input string) int {

    elves := strings.Split(input, "\n\n")
    result := 0

    for _, elve := range elves {
        calories := strings.Split(elve, "\n")
        sum := 0
        for _, calorie := range calories {
            calorie, _ := strconv.Atoi(calorie)
            sum += calorie
            if(sum > result) {
                result = sum
            }
        }

    }

    return result
}
