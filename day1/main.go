package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// problem: https://adventofcode.com/2021/day/1

	data, _ := os.ReadFile("input.txt")

	text := string(data)
	text = strings.TrimSpace(text)
	lines := (strings.Split(text, "\n"))
	intArr := make([]int, len(lines))
	for i := range intArr {
		intArr[i], _ = strconv.Atoi(lines[i])
	}

	// solution to part 1
	c := 0

	for i := 1; i < len(intArr); i++ {

		if intArr[i] > intArr[i-1] {
			c++
		}
	}

	// solution to part 2
	c2 := 0

	for i := 3; i < len(intArr); i++ {
		curr := intArr[i] + intArr[i-1] + intArr[i-2]
		prev := intArr[i-1] + intArr[i-2] + intArr[i-3]

		if curr > prev {
			c2++
		}
	}

	fmt.Println("part 1 answer:", c)
	fmt.Println("part 2 answer:", c2)

}
