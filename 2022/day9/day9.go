package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type location struct {
	vertical   int
	horizontal int
}
type rope struct {
	knots     []location
	tailStops map[string]struct{}
}

type instruction struct {
	direction string
	steps     int
}

func (r *rope) move() {
	for i := 1; i < len(r.knots); i++ {
		if r.checkKnotsSeparation(i) {
			r.knots[i].horizontal += signum(r.knots[i-1].horizontal - r.knots[i].horizontal)
			r.knots[i].vertical += signum(r.knots[i-1].vertical - r.knots[i].vertical)
		} else {
			break
		}
		fmt.Print()
	}
}

func (r *rope) checkKnotsSeparation(current int) bool {
	hSep := r.knots[current-1].horizontal - r.knots[current].horizontal
	vSep := r.knots[current-1].vertical - r.knots[current].vertical
	if math.Abs(float64(hSep)) > 1 {
		return true
	}
	if math.Abs(float64(vSep)) > 1 {
		return true
	}

	return false
}

func (r *rope) walkItOut(direction string, steps int) {
	for step := 0; step < steps; step++ {
		switch direction {
		case "L":
			r.knots[0].horizontal--
			r.move()
		case "R":
			r.knots[0].horizontal++
			r.move()
		case "U":
			r.knots[0].vertical++
			r.move()
		case "D":
			r.knots[0].vertical--
			r.move()
		}

		tailLoc := fmt.Sprintf("%d%d", r.knots[len(r.knots)-1].vertical, r.knots[len(r.knots)-1].horizontal)
		r.tailStops[tailLoc] = struct{}{}

	}

}
func getInput() []instruction {
	steps := strings.Split(input, "\n")
	returnData := []instruction{}
	for _, step := range steps {
		instructions := strings.Split(step, " ")
		instruction := instruction{}
		count, _ := strconv.Atoi(instructions[1])
		instruction.direction = instructions[0]
		instruction.steps = count
		returnData = append(returnData, instruction)
	}
	return returnData
}
func Solution() {
	data := getInput()
	// Cheesing the starting location because im too dumb to handle negative ints :(
	start := location{vertical: 1000, horizontal: 1000}

	// build ropes
	rope1, rope2 := rope{}, rope{}
	rope1.tailStops, rope2.tailStops = map[string]struct{}{}, map[string]struct{}{}
	rope1.knots = []location{start, start}
	rope2.knots = []location{start, start, start, start, start, start, start, start, start, start}

	for _, move := range data {
		rope1.walkItOut(move.direction, move.steps)
		rope2.walkItOut(move.direction, move.steps)
	}

	fmt.Printf("Day9 - part 1: %d\n", len(rope1.tailStops))
	fmt.Printf("Day9 - part 2: %d\n", len(rope2.tailStops))

}

func signum(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
