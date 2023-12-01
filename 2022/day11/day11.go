package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type test func(int) string
type operation func(int) int

type monkey struct {
	items           []int
	test            test
	operation       operation
	testVal         int
	inspectionCount int
}

func buildMonkeys(m map[string]*monkey) {
	sampleMod := strings.ReplaceAll(input, ":", "")
	for _, section := range strings.Split(sampleMod, "\n\n") {
		lines := strings.Split(section, "\n")
		monkeyName := strings.Split(lines[0], " ")[1]
		currMonkey := &monkey{}
		// Starting items
		startItemsInts := []int{}
		startingItems := strings.Split(strings.ReplaceAll(lines[1], "Starting items ", ""), ", ")
		for _, item := range startingItems {
			itemToInt, _ := strconv.Atoi(item)
			startItemsInts = append(startItemsInts, itemToInt)
		}
		// operation
		opslice := strings.Split(strings.ReplaceAll(lines[2], "Operation new = ", ""), " ")
		val := 0
		if opslice[2] != "old" {
			valInt, _ := strconv.Atoi(opslice[2])
			val = valInt
		}
		switch opslice[1] {
		case "+":
			if opslice[2] == "old" {
				currMonkey.operation = func(a int) int { return a + a }
			} else {
				currMonkey.operation = func(a int) int { return a + val }
			}
		case "*":
			if opslice[2] == "old" {
				currMonkey.operation = func(a int) int { return a * a }
			} else {
				currMonkey.operation = func(a int) int { return a * val }
			}
		}

		// test
		testNumString := strings.ReplaceAll(lines[3], "Test divisible by ", "")
		testNumInt, _ := strconv.Atoi(testNumString)
		trueDest := strings.ReplaceAll(lines[4], "If true throw to monkey ", "")
		trueDest = strings.ReplaceAll(trueDest, " ", "")
		falseDest := strings.ReplaceAll(lines[5], "If false throw to monkey ", "")
		falseDest = strings.ReplaceAll(falseDest, " ", "")
		currMonkey.testVal = testNumInt
		currMonkey.test = func(a int) string {
			if a%testNumInt == 0 {
				return trueDest
			} else {
				return falseDest
			}
		}

		currMonkey.items = startItemsInts
		m[monkeyName] = currMonkey
	}
}

func throw(m map[string]*monkey, to, from string, item int) {
	m[to].items = append(m[to].items, item)
	m[from].items = m[from].items[1:]
}

func runRound(monkeys map[string]*monkey, mod int) {
	for i := 0; i < len(monkeys); i++ {
		name := fmt.Sprint(i)
		monkey := monkeys[name]
		for _, item := range monkey.items {
			monkey.inspectionCount++
			worryLevel := item % mod
			// worryLevel := item
			worryLevel = monkey.operation(worryLevel)
			// worryLevel = worryLevel / 3
			destMonkey := monkey.test(worryLevel)
			throw(monkeys, destMonkey, name, worryLevel)
		}
	}
}
func Solution() {
	monkeys := map[string]*monkey{}
	buildMonkeys(monkeys)
	round := 0

	// I had to look this up, I didn't come up with this myself.
	mod := 1
	for _, m := range monkeys {
		mod *= m.testVal
	}
	for round < 10000 {
		runRound(monkeys, mod)
		round++
	}
	mostIterations := []int{}
	for _, m := range monkeys {
		mostIterations = append(mostIterations, m.inspectionCount)
	}
	fmt.Println("")
	sort.Slice(mostIterations, func(i, j int) bool {
		return mostIterations[j] < mostIterations[i]
	})
	total := mostIterations[0] * mostIterations[1]
	fmt.Printf("Day 11 - part1: %d\n", total)

}

//2637590098
