package main

import (
	"fmt"
	"math"
)

func main() {
	l1, l2, rMap := parseAndSort(data1)

	total := 0
	total2 := 0
	for i := 0; i < len(l1); i++ {
		// Part 1
		dif := int(math.Abs(float64(l1[i]) - float64(l2[i])))
		total += dif

		// Part 2
		cur := l1[i]
		count, ok := rMap[cur]
		if ok {
			total2 += cur * count
		}
	}
	fmt.Println(total)
	fmt.Println(total2)
}
