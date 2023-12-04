package main

import (
	"fmt"
	"strings"
)

var sample = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func main() {
	lines := strings.Split(input, "\n")
	total := 0
	total2 := 0

	// make a preallocated slice we can so we can reference any index immediately
	cards := make([]int, len(lines))

	for i, line := range lines {
		count := 0
		numberOfWinners := 0

		//each card will have at least 1 copy
		cards[i]++
		currentNumberOfCopies := cards[i]

		t := strings.Split(strings.ReplaceAll(line, ":  ", ": "), ": ")[1]
		card := strings.Split(t, " | ")

		winningNums := strings.Split(strings.ReplaceAll(card[0], "  ", " "), " ")
		myNums := strings.Split(strings.ReplaceAll(card[1], "  ", " "), " ")

		for _, num := range myNums {
			if isWinningNumber(winningNums, num) {
				//part1
				if count == 0 {
					count++
				} else {
					count *= 2
				}
				//part2
				numberOfWinners++
				if i+numberOfWinners < len(lines) {
					// update copy count at appropriate index in our slice
					cards[i+numberOfWinners] += currentNumberOfCopies
				}

			}
		}

		total += count

	}

	fmt.Printf("part 1 total: %v\n", total)
	for _, c := range cards {
		total2 += c
	}
	fmt.Printf("part 2 total: %v\n", total2)
}

func isWinningNumber(winners []string, test string) bool {
	for _, winner := range winners {
		if test == winner {
			return true
		}
	}
	return false
}
