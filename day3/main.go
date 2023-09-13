package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// problem: https://adventofcode.com/2021/day/3
	data, _ := os.ReadFile("sample.txt")
	text := strings.TrimSpace(string(data))
	lines := (strings.Split(text, "\n"))

	// part 1:
	gamma := make([]string, 0)
	eps := make([]string, 0)

	for i := 0; i < len(lines[0]); i++ {
		zeros := 0
		ones := 0
		for _, line := range lines {
			rune := string(line[i])
			if rune == "0" {
				zeros++
			} else {
				ones++
			}
		}
		if zeros > ones {
			gamma = append(gamma, "0")
			eps = append(eps, "1")
		} else {
			gamma = append(gamma, "1")
			eps = append(eps, "0")
		}
	}

	gammaDecimal, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 0)
	epsilonDecimal, _ := strconv.ParseInt(strings.Join(eps, ""), 2, 0)
	powerConsumption := gammaDecimal * epsilonDecimal

	fmt.Println("part 1 answer:", powerConsumption)

	// part 2
	oxygenCriteria := lines
	co2Criteria := lines
	var oxygenGeneratorRating string
	var co2ScrubberRating string

	// filter criteria for oxygen
	for i := 0; i < len(lines[0]); i++ {
		zeros := 0
		ones := 0
		for _, line := range oxygenCriteria {
			rune := string(line[i])
			if rune == "0" {
				zeros++
			} else {
				ones++
			}
		}
		if ones >= zeros {
			oxygenCriteria = discardContaining(oxygenCriteria, i, 0)
		} else {
			oxygenCriteria = discardContaining(oxygenCriteria, i, 1)
		}

		if len(oxygenCriteria) == 1 {
			oxygenGeneratorRating = oxygenCriteria[0]
			break
		}
	}

	// filter criteria for co2
	for i := 0; i < len(lines[0]); i++ {
		zeros := 0
		ones := 0
		for _, line := range co2Criteria {
			rune := string(line[i])
			if rune == "0" {
				zeros++
			} else {
				ones++
			}
		}
		if zeros <= ones {
			co2Criteria = discardContaining(co2Criteria, i, 1)
		} else {
			co2Criteria = discardContaining(co2Criteria, i, 0)
		}

		if len(co2Criteria) == 1 {
			co2ScrubberRating = co2Criteria[0]
			break
		}
	}
	a, _ := strconv.ParseInt(oxygenGeneratorRating, 2, 0)
	b, _ := strconv.ParseInt(co2ScrubberRating, 2, 0)
	fmt.Println("part 2 answer:", a*b)
}

func discardContaining(slice []string, index int, discard int) []string {
	items := make([]string, 0, len(slice))
	for _, value := range slice {
		if discard == 1 {
			if string(value[index]) == "0" {
				items = append(items, value)
			}
		} else {
			if string(value[index]) == "1" {
				items = append(items, value)
			}
		}
	}
	return items
}
