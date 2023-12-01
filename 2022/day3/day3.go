package day3

import (
	"aoc/utils"
	"strings"
)

func Solution() (int, int) {
	data := utils.GetInputData(3)
	rucks := strings.Split(data, "\n")

	// part 1
	total := 0
	for i := 0; i < len(rucks); i++ {
		ruck := rucks[i]
		c1, c2 := ruck[:len(ruck)/2], ruck[len(ruck)/2:]
		for _, char := range c1 {
			if strings.Contains(c2, string(char)) {
				total += bitToInt(byte(char))
				break
			}
		}
	}

	// part 2
	total2 := 0
	for i := 0; i < len(rucks); i += 3 {
		e1, e2, e3 := rucks[i], rucks[i+1], rucks[i+2]
		for _, char := range e1 {
			if strings.Contains(e2, string(char)) && strings.Contains(e3, string(char)) {
				total2 += bitToInt(byte(char))
				break
			}
		}
	}
	return total, total2
}

func bitToInt(b byte) int {
	var num int

	if b >= 97 {
		num += int(b - 'a' + 1)
	} else {
		num += int(b - 'A' + 27)
	}
	return num
}
