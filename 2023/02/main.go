package main

import (
	"fmt"
	"strconv"
	"strings"
)

var sample = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

type game struct {
	id         int
	red        int
	blue       int
	green      int
	isPossible bool
}

func main() {
	lines := strings.Split(input, "\n")

	total := 0
	total2 := 0
	for _, line := range lines {
		sets := strings.Split(line, ": ")
		gameId := strings.Split(sets[0], " ")[1]

		rounds := strings.Split(sets[1], "; ")
		intId, _ := strconv.Atoi(gameId)
		game := game{id: intId, isPossible: true}

		for _, round := range rounds {
			cubeSelection := strings.Split(round, ", ")
			for _, selection := range cubeSelection {
				cubes := strings.Split(selection, " ")
				count, _ := strconv.Atoi(cubes[0])
				color := cubes[1]

				switch color {
				case "red":
					if count > game.red {
						game.red = count
					}
					if count > 12 {
						game.isPossible = false
					}
				case "blue":
					if count > game.blue {
						game.blue = count
					}
					if count > 14 {
						game.isPossible = false
					}
				case "green":
					if count > game.green {
						game.green = count
					}
					if count > 13 {
						game.isPossible = false
					}
				}
			}

		}

		// part1 total
		if game.isPossible {
			total += game.id
		}

		//part2 total
		total2 += game.red * game.blue * game.green

	}

	fmt.Printf("part1 total: %v\n", total)
	fmt.Printf("part1 total2: %v\n", total2)
}
