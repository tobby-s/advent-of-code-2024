package puzzles

import (
	"log"
	"strconv"
	"strings"

	"github.com/tobby-s/advent-of-code-2024/utils"
)

func D2P1() int {
	url := "https://adventofcode.com/2024/day/2/input"
	data, err := utils.LoadData(url)
	if err != nil {
		log.Fatal(err)
	}
	safeCount := 0
	for _, s := range data {
		ints := strings.Split(s, " ")
		dir := 0
		var prev *int
		safe := true
		if s != "" {
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
			if safe {
				safeCount += 1
			}
		}

	}
	return safeCount
}