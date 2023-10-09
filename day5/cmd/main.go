package main

import (
	"day5/internal/part1"
	"day5/internal/part2"
	"io"
	"log"
	"os"
)

func main() {
	log.SetPrefix("day5: ")
	log.SetFlags(0)

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Part 1: %s", part1.Result(string(input)))
	log.Printf("Part 2: %s", part2.Result(string(input)))
}
