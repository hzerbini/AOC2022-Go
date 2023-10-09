package main

import (
	"day6/internal/part1"
	"day6/internal/part2"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("day6: ")

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Parte 1: %d", part1.Result(string(input)))
	log.Printf("Parte 2: %d", part2.Result(string(input)))
}
