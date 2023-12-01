package day8

import (
	"fmt"
	"strconv"
	"strings"
)

type views struct {
	left   int
	right  int
	top    int
	bottom int
}

type coordinates struct {
	row int
	col int
}

func checkLeft(height int, row []int, position int, views *views) bool {
	if position == 0 {
		return true
	}
	if height <= row[position-1] {
		views.left += 1
		return false
	}
	views.left += 1
	return checkLeft(height, row, position-1, views)
}

func checkRight(height int, row []int, position int, views *views) bool {
	if position == len(row)-1 {
		return true
	}
	if height <= row[position+1] {
		views.right += 1
		return false
	}
	views.right += 1
	return checkRight(height, row, position+1, views)
}

func checkTop(height int, rows [][]int, position coordinates, views *views) bool {
	if position.row == 0 {
		return true
	}
	if height <= rows[position.row-1][position.col] {
		views.top += 1
		return false
	}
	position.row -= 1
	views.top += 1
	return checkTop(height, rows, position, views)
}

func checkBottom(height int, rows [][]int, position coordinates, views *views) bool {
	if position.row == len(rows)-1 {
		return true
	}
	if height <= rows[position.row+1][position.col] {
		views.bottom += 1
		return false
	}
	position.row += 1
	views.bottom += 1
	return checkBottom(height, rows, position, views)
}

func convertGridToInts() [][]int {
	rows := strings.Split(input, "\n")
	grid := [][]int{}

	for row := 0; row < len(rows); row++ {
		tempRow := []int{}
		for col := 0; col < len(rows[0]); col++ {
			val, _ := strconv.Atoi(string(rows[row][col]))
			tempRow = append(tempRow, val)
		}
		grid = append(grid, tempRow)
	}
	return grid
}
func Solution() {
	rows := convertGridToInts()

	visibleTrees := (len(rows)*2 - 4) + (len(rows[0]) * 2)
	theCrib := 0
	var left, right, top, bottom bool

	for row := 1; row < len(rows)-1; row++ {
		for col := 1; col < len(rows[0])-1; col++ {
			coordinates := coordinates{row: row, col: col}
			val := rows[row][col]
			views := &views{}

			left = checkLeft(val, rows[row], col, views)
			right = checkRight(val, rows[row], col, views)
			top = checkTop(val, rows, coordinates, views)
			bottom = checkBottom(val, rows, coordinates, views)

			// part2
			viewDistance := views.left * views.bottom * views.right * views.top
			if viewDistance > theCrib {
				theCrib = viewDistance
			}

			// part1
			if left {
				visibleTrees += 1
				continue
			}

			if right {
				visibleTrees += 1
				continue
			}

			if top {
				visibleTrees += 1
				continue
			}

			if bottom {
				visibleTrees += 1
				continue
			}

		}

	}

	fmt.Printf("The Crib: %d\n", theCrib)
	fmt.Printf("Visible Trees: %d\n", visibleTrees)
}
