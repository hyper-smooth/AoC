package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	reg := regexp.MustCompile(" {2,}")
	data := reg.ReplaceAllString(input, " ")
	data = strings.ReplaceAll(data, "Time:", "")
	data = strings.ReplaceAll(data, "Distance:", "")

	lines := strings.Split(data, "\n")
	times := strings.Split(strings.Trim(lines[0], " "), " ")
	distance := strings.Split(strings.Trim(lines[1], " "), " ")

	pt1 := calcWins(times, distance)

	times2 := strings.Split(strings.ReplaceAll(lines[0], " ", ""), " ")
	distnace2 := strings.Split(strings.ReplaceAll(lines[1], " ", ""), " ")

	pt2 := calcWins(times2, distnace2)

	fmt.Println(pt1)
	fmt.Println(pt2)
}

func calcWins(times, distance []string) int {
	rw := 1
	for i := 0; i < len(times); i++ {
		t := utils.StringToInt(times[i])
		d := utils.StringToInt(distance[i])

		var first, last int

		// find the first win and mark its index then stop
		for j := 0; j < t; j++ {
			l := j * (t - j)
			if l > d {
				first = j
				break
			}
		}

		// start from the end to find the last win, mark its index + 1 and stop
		for hold := t; hold > 0; hold-- {
			l := hold * (t - hold)

			if l > d {
				last = hold + 1
				break
			}
		}

		// update number of wins by subtracting first from last to get count
		rw *= last - first
	}
	return rw
}
