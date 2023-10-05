package main

import (
	"day3/internal/part1"
	"day3/internal/part2"
	"io"
	"log"
	"os"
)

func main() {
	log.SetPrefix("day3: ")
	log.SetFlags(0)

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Part 1: %d", part1.Result(string(input)))
	log.Printf("Part 2: %d", part2.Result(string(input)))
}
