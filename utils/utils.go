package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

/*
Get input data by day.
@returns a multi line string as given by AoC
*/
func GetInputData(day int) string {
	client := http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)
	session := os.Getenv("AOC_SESSION")
	req, err := http.NewRequest("GET", url, nil)

	// The errors should really be returned and handled else where. I am just doing this to speed things along
	ErrorCheck(err)

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Cookie":       {fmt.Sprintf("session=%s", session)},
	}

	res, err := client.Do(req)
	ErrorCheck(err)

	body, err := io.ReadAll(res.Body)
	ErrorCheck(err)

	return strings.TrimSuffix(string(body), "\n")
}

func StringToInt(s string) int {
	num := 0
	bits := []byte(s)

	for _, bit := range bits {
		num = num*10 + (int(bit) - '0')
	}

	return num
}

func InRange(i, min, max int) bool {
	return (i >= min) && (i <= max)
}

func LetterToInt(b byte) int {
	var num int

	if b >= 97 {
		num += int(b - 'a' + 1)
	} else {
		num += int(b - 'A' + 27)
	}
	return num
}
