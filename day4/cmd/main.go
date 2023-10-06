package main

import (
	"day4/internal/part1"
	"day4/internal/part2"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
    log.SetFlags(0)
    log.SetPrefix("day4: ")
    input, err := io.ReadAll(os.Stdin)
    if err != nil  {
        log.Fatal(err)
    }

    fmt.Printf("Part1: %d\n", part1.Result(string(input)))
    fmt.Printf("Part2: %d\n", part2.Result(string(input)))
}
