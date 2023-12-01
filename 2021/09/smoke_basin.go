package smokebasin

import (
	"aoc/2021/inputs"
	"fmt"
	"strconv"
	"strings"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
func splitInputs() [][]int {
	data := inputs.GetInputData(9)
	mainSlice := strings.Split(data, "\n")
	finalSlice := [][]int{}
	for i := 0; i < len(mainSlice); i++ {
		tempSliceStr := strings.Split(mainSlice[i], "")
		tempSliceInt := []int{}

		for j := 0; j < len(tempSliceStr); j++ {
			tempInt, err := strconv.Atoi(tempSliceStr[j])
			errorCheck(err)
			tempSliceInt = append(tempSliceInt, tempInt)
		}
		finalSlice = append(finalSlice, tempSliceInt)

	}
	return finalSlice
}

func Solution1() {
	tunnelMap := splitInputs()
	lowPointsList := []int{}

	for i := 0; i < len(tunnelMap); i++ {
		for j := 0; j < len(tunnelMap[i]); j++ {
			currentPosition := tunnelMap[i][j]

			// compare to top val
			if i != 0 {
				if currentPosition >= tunnelMap[i-1][j] {
					continue
				}
			}

			// compare to bottom val
			if i < len(tunnelMap)-1 {
				if currentPosition >= tunnelMap[i+1][j] {
					continue
				}
			}

			//compare to left value
			if j != 0 {
				if currentPosition >= tunnelMap[i][j-1] {
					continue
				}
			}

			//compare to right value
			if j < len(tunnelMap[i])-1 {
				if currentPosition >= tunnelMap[i][j+1] {
					continue
				}
			}
			lowPointsList = append(lowPointsList, currentPosition)
		}
	}
	totalRiskScore := 0
	for i := 0; i < len(lowPointsList); i++ {
		risk := lowPointsList[i] + 1
		totalRiskScore += risk
	}
	fmt.Printf("Day 9 - part 1 answer: %d\n", totalRiskScore)
}
