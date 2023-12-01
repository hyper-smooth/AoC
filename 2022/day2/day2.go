package day2

import (
	"aoc/utils"
	"strings"
)

func Solution() (int, int) {
	data := utils.GetInputData(2)
	total := 0
	total2 := 0

	enum := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

	for _, match := range strings.Split(data, "\n") {
		selections := strings.Split(match, " ")
		op, mine := enum[selections[0]], enum[selections[1]]

		// Part 1
		outcome := (mine - op) % 3
		switch outcome {
		case 0:
			total += 3
		case 1:
			total += 6
		case -2:
			total += 6
		}
		total += mine
		// Part 2
		win := map[int]int{1: 2, 2: 3, 3: 1}
		lose := map[int]int{1: 3, 2: 1, 3: 2}

		switch mine {
		case 1:
			total2 += lose[op]
		case 2:
			total2 += op
			total2 += 3
		case 3:
			total2 += win[op]
			total2 += 6
		}
	}

	return total, total2
}
