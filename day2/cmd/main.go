package main

import (
	"day2/internal/part1"
	"day2/internal/part2"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
    log.SetPrefix("day2: ")
    log.SetFlags(0)
    input, err := io.ReadAll(os.Stdin)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Part1: %d\n", part1.Result(string(input)))
    fmt.Printf("Part2: %d\n", part2.Result(string(input)))
}
