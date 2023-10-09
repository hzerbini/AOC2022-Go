package part2

func Result(input string) int {
	i := 14
	found := false

	for i < len(input) && found == false {
		found = true
		substr := input[i-14 : i]
		c1 := 0
		count := 0

		for c1 < len(substr) && count < 2 {
			count = 0
			c2 := 0

			for c2 < len(substr) && count < 2 {
				if substr[c1] == substr[c2] {
					count++
				}
				c2++
			}

			if count > 1 {
				found = false
			}
			c1++
		}

		i++
	}

	if found {
		return i - 1
	} else {
		return -1
	}

}
