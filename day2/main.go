package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// problem: https://adventofcode.com/2021/day/2
	data, _ := os.ReadFile("input.txt")
	text := string(data)
	text = strings.TrimSpace(text)
	lines := (strings.Split(text, "\n"))

	// part 1

	x := 0
	y := 0

	for _, line := range lines {
		commands := strings.Split(line, " ")
		dir, strAmount := commands[0], commands[1]
		amount, _ := strconv.Atoi(strAmount)
		if dir == "down" {
			y += amount
		} else if dir == "up" {
			y -= amount
		} else if dir == "forward" {
			x += amount
		}
	}

	ans := x * y

	// part 2

	x, y, aim := 0, 0, 0

	for _, line := range lines {
		commands := strings.Split(line, " ")
		dir := commands[0]
		amount, _ := strconv.Atoi(commands[1])

		if dir == "down" {
			aim += amount
		} else if dir == "up" {
			aim -= amount
		} else if dir == "forward" {
			x += amount
			y += amount * aim
		}
	}

	ans2 := x * y

	fmt.Println(ans)
	fmt.Println(ans2)
}
