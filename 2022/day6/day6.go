package day6

import (
	"aoc/utils"
)

func Solution1() int {
	sigStart := 0
	data := utils.GetInputData(6)
	tempString := data[:4]
	for i := 3; i < len(data); i++ {
		m := map[string]int{}
		for _, char := range tempString {
			m[string(char)] += 1
		}

		if len(m) == 4 {
			sigStart = i
			break
		}
		s := string(data[i])
		tempString = tempString[1:] + s
	}
	return sigStart
}

func Solution2() int {
	msgStart := 0
	data := utils.GetInputData(6)
	tempString := data[:14]

	for i := 14; i < len(data); i++ {
		m := map[string]int{}
		for _, char := range tempString {
			m[string(char)] += 1
		}

		if len(m) == 14 {
			msgStart = i
			break
		}
		s := string(data[i])
		tempString = tempString[1:] + s

	}
	return msgStart
}
func Start() (int, int) {
	return Solution1(), Solution2()
}
