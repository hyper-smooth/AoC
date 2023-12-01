package day10

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	cycle        int
	register     int
	tracker      map[int]int
	total        int
	linePosition int
	crt          string
	eventList    []int
}

func (c *cpu) incrementCycle() {
	c.cycle++
	c.record()
}

func (c *cpu) updateRegister(a int) {
	c.incrementCycle()
	c.draw()
	c.incrementCycle()
	c.register += a
}

func (c *cpu) record() {
	if c.cycle%20 == 0 {
		c.tracker[c.cycle] = c.register
	}
}

func (c *cpu) calcTotal() {
	for _, event := range c.eventList {
		c.total += event * c.tracker[event]
	}
}

func (c *cpu) draw() {
	if utils.InRange(c.linePosition, c.register-1, c.register+1) {
		c.crt += "#"
	} else {
		c.crt += "."
	}
	c.linePosition++

	//start drawing next line
	if c.linePosition == 40 {
		c.linePosition = 0
		c.crt += "\n"
	}

}

func Solution() {
	instructionSet := strings.Split(utils.GetInputData(10), "\n")

	cpu := cpu{}
	cpu.tracker = map[int]int{}
	cpu.register = 1
	cpu.eventList = []int{20, 60, 100, 140, 180, 220} // cycle number to record

	for _, instruction := range instructionSet {
		cpu.draw()
		if instruction == "noop" {
			cpu.incrementCycle()
			continue
		}
		amount := strings.Split(instruction, " ")[1]
		a, _ := strconv.Atoi(amount)
		cpu.updateRegister(a)
	}

	cpu.calcTotal()

	fmt.Printf("Day 10 - part1: %d\n", cpu.total)
	fmt.Printf("Day 10 - part2:\n%s\n", cpu.crt)

}
