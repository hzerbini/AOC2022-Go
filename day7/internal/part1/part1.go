package part1

import (
	"fmt"
	"strconv"
	"strings"
)

func Result(input string) int {
	lines := strings.Split(input, "\n")
	sizes := parse(lines)
	result := 0

	for _, size := range sizes {
		if size < 100000 {
			result += size
		}
	}

	return result
}

func parse(lines []string) map[string]int {
	result := map[string]int{}
	pwd := []string{""}

	for i := 0; i < len(lines)-1; i++ {
		command, found := strings.CutPrefix(lines[i], "$ ")
		if found {
			c := strings.Fields(command)

			if c[0] == "cd" {
				if len(c) < 2 {
					panic(fmt.Errorf("Invalid command: %s", command))
				}

				if c[1] == ".." && len(pwd) > 1 {
					pwd = pwd[:len(pwd)-1]
				} else if c[1] == "/" {
					pwd = []string{""}
				} else if c[1] != ".." {
					pwd = append(pwd, c[1])
				}
			}
		} else {
			output := strings.Fields(lines[i])
			if len(output) < 2 {
				panic(fmt.Errorf("Invalid output: %s", lines[i]))
			}

			num, err := strconv.Atoi(output[0])

			if err == nil {
				for j := len(pwd); j > 0; j-- {
					result[strings.Join(pwd[:j], "/")] += num
				}
			}
		}
	}

	return result
}
