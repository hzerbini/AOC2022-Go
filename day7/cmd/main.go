package main

import (
	"io"
	"log"
	"os"

	"day7/internal/part1"
	"day7/internal/part2"
)

func main() {
	log.SetPrefix("day7: ")
	log.SetFlags(0)
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Part 1: %d", part1.Result(string(input)))
	log.Printf("Part 2: %d", part2.Result(string(input)))
}
