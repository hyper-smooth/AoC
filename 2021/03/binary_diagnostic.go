package binarydiagnostic

import (
	"strings"
)

func Part1() int {
	diagnostics := strings.Split(input, "\n")

	// Renamed the variables which maintain the "shape" of the binary input.
	n, m := len(diagnostics), len(diagnostics[0])

	// Prefer using math to maintain the binary values we build up
	// instead of leveraging strings and built-ins which parse them
	// to decimal.
	gamma := make([]byte, m)
	epsilon := make([]byte, m)

	for col := 0; col < m; col++ {

		// Whenever a map can be indexed by integers in consecutive
		// order, you can replace it with a slice (or array).
		//  - onOff[0] maintains the count of zero
		//  - onOff[1] maintains the count of one
		var onOff [2]int

		for row := 0; row < n; row++ {
			// We know that the characters are only '0' and '1' so
			// every character is a single byte and we can use
			// naive indexing.
			bits := []byte(diagnostics[row])

			// Subtracting the lowest character from the range of character
			// values is a trick to get the integer value. This only works
			// for ASCII characters since they're always sorted ascending.
			// https://stackoverflow.com/a/3195042
			bit := bits[col] - '0'
			onOff[bit]++
		}
		if onOff[1] > onOff[0] {
			gamma[col]++
		} else {
			epsilon[col]++
		}
	}

	// Returning the answer value to allow us to test this
	// on other inputs, or do benchmarking.

	return bitsToInt(gamma) * bitsToInt(epsilon)
}
func Part2() int {

	diagnostics := strings.Split(input, "\n")

	m := len(diagnostics[0])

	oxygneGeneratorContainer := diagnostics
	co2ScrubberContainer := diagnostics

	// Get Oxygen generator rating
	for col := 0; col < m; col++ {
		if len(oxygneGeneratorContainer) <= 1 {
			continue
		}
		var onOff [2]int

		onOffMap := make(map[byte][]string)

		for _, row := range oxygneGeneratorContainer {

			bits := []byte(row)

			bit := bits[col] - '0'

			onOff[bit]++
			onOffMap[bit] = append(onOffMap[bit], row)
		}
		if onOff[1] >= onOff[0] {
			oxygneGeneratorContainer = onOffMap[1]
		} else {
			oxygneGeneratorContainer = onOffMap[0]
		}
	}

	// Get CO2 scrubber rating
	// This loop could probably be combined with the one above
	for col := 0; col < m; col++ {
		if len(co2ScrubberContainer) <= 1 {
			continue
		}
		var onOff [2]int

		onOffMap := make(map[byte][]string)

		for _, row := range co2ScrubberContainer {
			bits := []byte(row)

			bit := bits[col] - '0'
			onOff[bit]++
			onOffMap[bit] = append(onOffMap[bit], row)
		}

		if onOff[1] < onOff[0] {
			co2ScrubberContainer = onOffMap[1]
		} else {
			co2ScrubberContainer = onOffMap[0]
		}
	}
	oxygenGeneratorRating := make([]byte, len(oxygneGeneratorContainer[0]))
	co2ScrubberRating := make([]byte, len(co2ScrubberContainer[0]))
	oxygenBits := []byte(oxygneGeneratorContainer[0])
	co2Bits := []byte(co2ScrubberContainer[0])

	for i := range co2Bits {
		co2Bit := co2Bits[i] - '0'
		oxygenBit := oxygenBits[i] - '0'

		oxygenGeneratorRating[i] = oxygenBit
		co2ScrubberRating[i] = co2Bit
	}

	return bitsToInt(oxygenGeneratorRating) * bitsToInt(co2ScrubberRating)
}

func bitsToInt(bits []byte) int {
	v := 0
	for i, b := range bits {
		// https://stackoverflow.com/a/23189744
		v += int(b) << (len(bits) - (i + 1))
	}
	return v
}
