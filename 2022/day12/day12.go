package day12

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type location struct {
	x int
	y int
}

type my struct {
	steps    int
	location location
}

func buildMap() ([][]int, [2]int) {
	grid := [][]int{}
	rows := strings.Split(sample, "\n")
	destLoc := [2]int{}
	for row := 0; row < len(rows); row++ {
		gridRow := []int{}
		for i, col := range rows[row] {
			letterAsInt := utils.LetterToInt(byte(col))
			if string(col) == "S" {
				letterAsInt = utils.LetterToInt('a')
			}
			if string(col) == "E" {
				destLoc[0], destLoc[1] = row, i
				letterAsInt = utils.LetterToInt('z')
			}

			gridRow = append(gridRow, letterAsInt)
		}
		grid = append(grid, gridRow)
	}
	return grid, destLoc
}
func Solution() {
	data, destLoc := buildMap()
	fmt.Println(data)
	fmt.Println(destLoc)
}
