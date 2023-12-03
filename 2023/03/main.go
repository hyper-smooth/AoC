package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var sample = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func main() {
	rows := strings.Split(input, "\n")

	total := 0
	total2 := 0
	gears := map[string][]string{}
	for rowIndex, row := range rows {
		scanner := ""
		isValid := false
		gearCords := ""
		for colIndex, col := range row {
			// if its  a number add to scanner and check validity
			if utils.InRange(int(col), 48, 57) {
				scanner += string(col)
				//check up
				if rowIndex > 0 {
					if isSymbol(rows[rowIndex-1][colIndex]) {
						gearCords = fmt.Sprintf("%v:%v", rowIndex-1, colIndex)
						isValid = true
					}

					if colIndex != 0 && isSymbol(rows[rowIndex-1][colIndex-1]) {
						gearCords = fmt.Sprintf("%v:%v", rowIndex-1, colIndex-1)
						isValid = true
					}
					if colIndex != len(rows)-1 && isSymbol(rows[rowIndex-1][colIndex+1]) {
						gearCords = fmt.Sprintf("%v:%v", rowIndex-1, colIndex+1)
						isValid = true
					}
				}
				//check left
				if colIndex != 0 && isSymbol(rows[rowIndex][colIndex-1]) {
					gearCords = fmt.Sprintf("%v:%v", rowIndex, colIndex-1)
					isValid = true
				}
				//check right
				if colIndex != len(rows)-1 && isSymbol(rows[rowIndex][colIndex+1]) {
					gearCords = fmt.Sprintf("%v:%v", rowIndex, colIndex+1)
					isValid = true
				}
				//down
				if rowIndex < len(rows)-1 {

					if isSymbol(rows[rowIndex+1][colIndex]) {
						gearCords = fmt.Sprintf("%v:%v", rowIndex+1, colIndex)
						isValid = true
					}

					if colIndex != 0 && isSymbol(rows[rowIndex+1][colIndex-1]) {
						gearCords = fmt.Sprintf("%v:%v", rowIndex+1, colIndex-1)
						isValid = true
					}
					if colIndex != len(rows)-1 && isSymbol(rows[rowIndex+1][colIndex+1]) {
						gearCords = fmt.Sprintf("%v:%v", rowIndex+1, colIndex+1)
						isValid = true
					}
				}

				if colIndex == len(rows)-1 {
					total += utils.StringToInt(scanner)
					upsertGears(gears, gearCords, scanner)
				}

				//go to next iteration
				continue

			}

			//if scanner has values and isValid then add to total
			if scanner != "" && isValid {
				total += utils.StringToInt(scanner)
			}

			// part 2
			if scanner != "" && gearCords != "" {
				upsertGears(gears, gearCords, scanner)
			}

			gearCords = ""
			scanner = ""
			isValid = false
		}

	}

	for gear, parts := range gears {
		if len(parts) == 2 {
			total2 += utils.StringToInt(parts[0]) * utils.StringToInt(parts[1])
		}
		fmt.Printf("gear: %v - parts: %v\n", gear, parts)
	}
	fmt.Printf("part 1 total: %v\n", total)
	fmt.Printf("part 2 total: %v\n", total2)
}

// comment this out if part 1
func isSymbol(x byte) bool {

	return string(x) == "*"

}

//comment this out if part 2
// func isSymbol(x byte) bool {

// 	if utils.InRange(int(x), 48, 57) {
// 		return false
// 	}
// 	//if .
// 	if x == 46 {
// 		return false
// 	}

// 	return true
// }

func upsertGears(gears map[string][]string, coords string, value string) {

	parts, ok := gears[coords]

	if !ok {
		gears[coords] = []string{value}
		return
	}

	parts = append(parts, value)

	gears[coords] = parts
}
