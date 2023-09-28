package main

import (
	"fmt"
	"hzerbini/day1/internal/part1"
	"hzerbini/day1/internal/part2"
    "log"
    "os"
    "io"
)

func main() {
    log.SetPrefix("day1: ")
    log.SetFlags(0)

    bytes, err := io.ReadAll(os.Stdin)
    if err != nil {
        log.Fatal(err)
    }

    input := string(bytes)

    fmt.Printf("Part1: %d\n", part1.Result(input))
    fmt.Printf("Part2: %d\n", part2.Result(input))
}
