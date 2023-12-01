package day5

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type move struct {
	count int
	from  int
	to    int
}

func formatData() ([9][]string, []move) {
	fStacks := strings.ReplaceAll(utils.GetInputData(5), "    [", "--- [")
	fStacks = strings.ReplaceAll(fStacks, "]    ", "] ---")
	fStacks = strings.ReplaceAll(fStacks, "     ", " --- ")
	fStacks = strings.ReplaceAll(fStacks, "    ", " ---")

	data := strings.Split(fStacks, "\n\n")
	stacks, moves := strings.Split(data[0], "\n"), strings.Split(data[1], "\n")

	instructions := []move{}
	stacksSlice := [9][]string{}
	stacks = stacks[:len(stacks)-1]

	for _, rows := range stacks {
		cols := strings.Split(rows, " ")

		for i, col := range cols {
			if col == "---" {
				continue
			}
			stacksSlice[i] = append(stacksSlice[i], col)
		}

	}
	for _, m := range moves {
		steps := strings.Split(m, " ")
		temp := move{}
		temp.count, _ = strconv.Atoi(steps[1])
		temp.from, _ = strconv.Atoi(steps[3])
		temp.to, _ = strconv.Atoi(steps[5])
		instructions = append(instructions, temp)
	}
	return stacksSlice, instructions
}
func Solution() {
	stackSlice, instructions := formatData()

	// Part 1
	for _, instruction := range instructions {
		for count := 0; count < instruction.count; count++ {
			crate := stackSlice[instruction.from-1][0]
			stackSlice[instruction.to-1] = append([]string{crate}, stackSlice[instruction.to-1]...)
			stackSlice[instruction.from-1] = stackSlice[instruction.from-1][1:]
		}
	}

	for _, t := range stackSlice {
		fmt.Print(t[0])
	}

	fmt.Println("---")

	// Part 2
	stackSlice, instructions = formatData()
	for _, instruction := range instructions {

		crates := make([]string, len(stackSlice[instruction.from-1][:instruction.count]))
		copy(crates, stackSlice[instruction.from-1][:instruction.count])
		stackSlice[instruction.to-1] = append(crates, stackSlice[instruction.to-1]...)
		stackSlice[instruction.from-1] = stackSlice[instruction.from-1][instruction.count:]
	}
	for _, t := range stackSlice {
		fmt.Print(t[0])
	}

}
