package main

import (
	"fmt"
	"strconv"
	"strings"
)

var sample = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func main() {
	part1()
	part2()
}

func part2() {
	lines := strings.Split(input, "\n")

	m := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	values := []int{}
	for _, line := range lines {
		bytes := []byte(line)

		var first, last, scanner string

		for _, b := range bytes {
			num := ""
			scanner += string(b)

			// If the byte is a numeral go ahead and reset scanner and set our number
			if b > 47 && b < 58 {
				num = string(b)
				scanner = ""
			}

			// Since the scanner may have started with a letter not part of the word
			// we need to check for sub strings of the scanner
			scanner = checkScannerForSubString(scanner)
			if dig, ok := m[scanner]; ok {
				num = dig
				//some string numbers share letters ex eightwo
				// we need to reset the scanner to the last letter if a match is found
				scanner = string(b)
			}

			if num == "" {
				continue
			}

			if first == "" {
				first = num
			}
			last = num
		}
		v := first + last
		fmt.Println(v)
		vInt, _ := strconv.Atoi(v)
		values = append(values, vInt)

	}

	total := 0
	for _, val := range values {
		total += val
	}
	fmt.Printf("total part 2: %v\n", total)
}

func part1() {
	lines := strings.Split(input, "\n")
	total := finder(lines)
	fmt.Printf("total part 1: %v\n", total)
}

func finder(lines []string) int {
	values := []int{}
	for _, line := range lines {
		bytes := []byte(line)

		var first, last string
		for _, b := range bytes {
			if b < 47 || b > 57 {
				continue
			}
			if first == "" {
				first = string(b)
			}
			last = string(b)
		}
		v := first + last

		vInt, _ := strconv.Atoi(v)
		values = append(values, vInt)

	}

	total := 0
	for _, val := range values {
		total += val
	}

	return total
}

func checkScannerForSubString(s string) string {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, word := range words {
		if strings.Contains(s, word) {
			return word
		}
	}
	return s
}
