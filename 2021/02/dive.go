package dive

import (
	"aoc/2021/inputs"
	"fmt"
	"strconv"
	"strings"
)

type submarine struct {
	location int
	depth    int
	aim      int
}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func Start() {
	sub := submarine{location: 0, depth: 0}
	subAdvanced := submarine{}
	directionsList := strings.Split(inputs.GetInputData(2), "\n")

	for _, directions := range directionsList {
		direction := strings.Split(directions, " ")
		amount, err := strconv.Atoi(direction[1])
		errorCheck(err)

		if direction[0] == "forward" {
			sub.location += amount
			subAdvanced.location += amount
			subAdvanced.depth += amount * subAdvanced.aim
		}

		if direction[0] == "down" {
			sub.depth += amount
			subAdvanced.aim += amount
		}
		if direction[0] == "up" {
			if amount > sub.depth {
				sub.depth = 0
			} else {
				sub.depth -= amount
			}

			if amount > subAdvanced.aim {
				subAdvanced.aim = 0
			} else {
				subAdvanced.aim -= amount
			}
		}
	}
	result := sub.depth * sub.location
	resultAdvanced := subAdvanced.depth * subAdvanced.location
	fmt.Printf("Day 2 - part 1 answer: %d\n", result)
	fmt.Printf("Day 2 - part 2 answer: %d\n", resultAdvanced)
}
