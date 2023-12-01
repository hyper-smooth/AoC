package day1

import (
	"aoc/utils"
	"sort"
	"strconv"
	"strings"
)

func Solution() (int, int) {
	payLoads := []int{}
	total := 0
	elves := strings.Split(input, "\n\n")

	for _, elf := range elves {
		caloryLoad := 0
		cals := strings.Split(elf, "\n")
		for _, cal := range cals {
			n, err := strconv.Atoi(cal)
			utils.ErrorCheck(err)
			caloryLoad += n
		}
		payLoads = append(payLoads, caloryLoad)

	}
	sort.Slice(payLoads, func(i, j int) bool {
		return payLoads[i] > payLoads[j]
	})

	for _, elf := range payLoads[:3] {
		total += elf
	}
	return payLoads[0], total
}
