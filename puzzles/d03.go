package puzzles

import (
	"regexp"
	"strconv"

	"github.com/tobby-s/advent-of-code-2024/utils"
)

func D3P1() int {
	url := "https://adventofcode.com/2024/day/3/input"
	data := utils.LoadData(url)
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	sum := 0
	for _, instructions := range data {
		matches := regex.FindAllStringSubmatch(instructions,-1)
		for _, m := range matches {
			num1, _ := strconv.Atoi(m[1])
			num2, _ := strconv.Atoi(m[2])
			sum += num1 * num2
		}
	}
	return sum
}

func D3P2() int {
	return 0
}