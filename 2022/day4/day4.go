package day4

import (
	"strconv"
	"strings"
)

type section struct {
	min int
	max int
}

func Solution() (int, int) {
	pairs := strings.Split(input, "\n")
	count := 0
	count2 := 0
	for _, pair := range pairs {
		elf1, elf2 := section{}, section{}

		sections := strings.Split(pair, ",")
		sector1 := strings.Split(sections[0], "-")
		sector2 := strings.Split(sections[1], "-")
		elf1.min, _ = strconv.Atoi(sector1[0])
		elf1.max, _ = strconv.Atoi(sector1[1])
		elf2.min, _ = strconv.Atoi(sector2[0])
		elf2.max, _ = strconv.Atoi(sector2[1])

		// part 1
		if elf1.min >= elf2.min && elf1.max <= elf2.max {
			count += 1
		} else if elf2.min >= elf1.min && elf2.max <= elf1.max {
			count += 1
		}

		// part 2
		if elf1.min >= elf2.min && elf1.min <= elf2.max {
			count2 += 1
		} else if elf2.min >= elf1.min && elf2.min <= elf1.max {
			count2 += 1
		}

	}
	return count, count2
}
