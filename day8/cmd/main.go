package main

import (
	"io"
	"log"
	"os"

	"day8/internal/part1"
	"day8/internal/part2"
)

func main() {
	log.SetPrefix("Day 8: ")
	log.SetFlags(0)
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Part1: %d\n", part1.Result(string(input)))
	log.Printf("Part2: %d\n", part2.Result(string(input)))
}
