package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var seedsSample = `79 14 55 13`
var sample = `seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func main() {
	lines := strings.Split(input, "\n\n")
	seedList := seedsP2()
	lowest := -1
	mappers := buildMaps(lines)
	// prep our seeds list which is an array of seed arrays
	for _, s := range seedList {
		currentSeed := []int{s}
		for i, m := range mappers {
			current := currentSeed[i]
			next := current
			for _, mp := range m {
				if current >= mp.start && current < mp.start+mp.r {
					next += mp.dest - mp.start
				}
			}

			currentSeed = append(currentSeed, next)
		}

		if currentSeed[len(currentSeed)-1] < lowest || lowest == -1 {
			lowest = currentSeed[len(currentSeed)-1]
		}
	}

	fmt.Printf("lowest location: %v\n", lowest)
}

type mapper struct {
	start int
	dest  int
	r     int
}

func buildMaps(lines []string) [][]mapper {
	rVal := [][]mapper{}
	for _, line := range lines {
		m := strings.Split(strings.ReplaceAll(line, " map", ""), ":\n")
		inner := []mapper{}
		for _, t := range strings.Split(m[1], "\n") {
			mm := mapper{}
			vals := strings.Split(t, " ")
			mm.dest = utils.StringToInt(vals[0])
			mm.start = utils.StringToInt(vals[1])
			mm.r = utils.StringToInt(vals[2])
			inner = append(inner, mm)

		}
		rVal = append(rVal, inner)
	}

	return rVal
}
func seedsP2() []int {
	ms := strings.Split(mySeeds, " ")
	//Before
	// newSeeds := []int{}
	//After
	newSeeds := make([]int, len(ms))

	for i := 0; i < len(ms); i += 2 {
		start := utils.StringToInt(ms[i])
		end := utils.StringToInt(ms[i+1])
		for j := start; j < start+end; j++ {

			// I was originally appending to an array which caused the runtime of the entire program to be around 3 minutes
			// newSeeds = append(newSeeds, j)

			// preallocating the slice and inserting into index reduced runtime to 500ms
			newSeeds[i] = j
		}
	}

	return newSeeds
}

// leaving this here as it works but you need a super computer to run it. it crushed my poor little macbook
type seed struct {
	id         int
	soil       int
	fertilizer int
	water      int
	light      int
	temp       int
	humidity   int
	location   int
}

func part1V1() {

	lines := strings.Split(input, "\n\n")
	seeds := []seed{}
	maps := map[string]map[int]int{}
	for _, line := range lines {
		m := strings.Split(strings.ReplaceAll(line, " map", ""), ":\n")
		ranges := map[int]int{}
		maps[m[0]] = ranges

		for _, t := range strings.Split(m[1], "\n") {
			vals := strings.Split(t, " ")
			dest := utils.StringToInt(vals[0])
			start := utils.StringToInt(vals[1])

			for i := 0; i < utils.StringToInt(vals[2]); i++ {
				ranges[start] = dest
				start++
				dest++
			}

		}

	}

	for _, s := range strings.Split(mySeeds, " ") {
		id := utils.StringToInt(s)
		newSeed := seed{id: id}

		newSeed.soil = checkMap(maps["seed-to-soil"], newSeed.id)
		newSeed.fertilizer = checkMap(maps["soil-to-fertilizer"], newSeed.soil)
		newSeed.water = checkMap(maps["fertilizer-to-water"], newSeed.fertilizer)
		newSeed.light = checkMap(maps["water-to-light"], newSeed.water)
		newSeed.temp = checkMap(maps["light-to-temperature"], newSeed.light)
		newSeed.humidity = checkMap(maps["temperature-to-humidity"], newSeed.temp)
		newSeed.location = checkMap(maps["humidity-to-location"], newSeed.humidity)
		seeds = append(seeds, newSeed)
	}

	lowest := seeds[0].location

	for _, s := range seeds {
		if s.location < lowest {
			lowest = s.location
		}
	}

	fmt.Printf("lowest location: %v\n", lowest)
}

func checkMap(m map[int]int, i int) int {
	if val, ok := m[i]; ok {
		return val
	}
	return i
}
