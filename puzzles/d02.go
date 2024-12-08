package puzzles

import (
	"log"
	"strconv"
	"strings"

	"github.com/tobby-s/advent-of-code-2024/utils"

)

func D2P1() int {
	url := "https://adventofcode.com/2024/day/2/input"
	data := utils.LoadData(url)
	safeCount := 0
	for _, s := range data {
		ints := strings.Split(s, " ")
		if s != "" {
			if isSafe(ints) {
				safeCount += 1
			}
		}

	}
	return safeCount
}

func D2P2() int {
	url := "https://adventofcode.com/2024/day/2/input"
	data := utils.LoadData(url)
	safeCount := 0 
	for _, s := range data {
		ints := strings.Split(s, " ")
		dampenedSafe := false
		if s != "" {
			if isSafe(ints) {
				safeCount += 1
				continue
			}
			for i, _ := range ints {
				ints2 := make([]string,0)
				ints2 = append(ints2, ints[:i]...)
				ints2 = append(ints2,ints[i+1:]...)
				if isSafe(ints2){
					dampenedSafe = true
					break
				}
			}
		}
		if dampenedSafe {
			safeCount += 1
		}
	}
	return safeCount
}

func isSafe(ints []string) bool {
	dir := 0
	var prev *int
	safe := true
	for _, i := range ints {
		num, _ := strconv.Atoi(i)
		if dir == 0 && prev != nil {
			if num == *prev {
				safe = false
				break 
			} else if num < *prev {
				dir = -1
			} else if num > *prev {
				dir = 1
			}
		}
		if prev != nil {
			diff := dir * (num - *prev)
			if diff < 1 || diff > 3 {
				safe = false
				break
			}
		}
		prev = &num
	}
	return safe
}