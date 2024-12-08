package puzzles

import (
	"fmt"
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
	url := "https://adventofcode.com/2024/day/3/input"
	data := utils.LoadData(url)
	regex := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	sum := 0
	enabled := true
	for _, instructions := range data {
		matches := regex.FindAllStringSubmatch(instructions,-1)
		for _, m := range matches {
			switch {
			case m[0] == "do()":
				enabled = true
			case m[0] == "don't()":
				enabled = false
			case len(m) >= 4:
				num1, _ := strconv.Atoi(m[2])
				num2, _ := strconv.Atoi(m[3])
				if enabled {
					sum += num1 * num2
				}
			}
		}
	}
	return sum
}