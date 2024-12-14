package puzzles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tobby-s/advent-of-code-2024/utils"
)

func D11P1() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/11/input")
	initial := strings.Split(data[0], " ")
	for i := 0; i < 25; i++ {
		initial = blink(initial)
	}

	return len(initial)
}

func D11P2() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/11/input")
	initial := strings.Split(data[0], " ")
	cache := map[xy]int{}
	for _, stone := range initial {
		n, _ := strconv.Atoi(stone)
		result += noStones(cache, n, 75)
	}
	return
}

func blink(stones []string) []string {
	newstones := make([]string, 0)
	for _, stone := range stones {
		num, _ := strconv.Atoi(stone)
		actualnum := fmt.Sprintf("%d", num)
		l := len(actualnum)
		if num == 0 {
			newstones = append(newstones, "1")
		} else if l%2 == 0 {
			newstones = append(newstones, actualnum[:(l/2)], actualnum[(l/2):])
		} else {
			newstones = append(newstones, fmt.Sprintf("%d", num*2024))
		}
	}
	return newstones
}

func noStones(cache map[xy]int, stone int, blinks int) int {
	if blinks == 0 {
		cache[xy{stone, 0}] = 1
		return 1
	}
	if n, ok := cache[xy{stone, blinks}]; ok {
		return n
	}
	strnum := fmt.Sprintf("%d", stone)
	l := len(strnum)
	if stone == 0 {
		cache[xy{stone, blinks}] = noStones(cache, 1, blinks-1)
	} else if l%2 == 0 {
		lnum, _ := strconv.Atoi(strnum[:(l / 2)])
		rnum, _ := strconv.Atoi(strnum[(l / 2):])
		cache[xy{stone, blinks}] = noStones(cache, lnum, blinks-1) + noStones(cache, rnum, blinks-1)
	} else {
		cache[xy{stone, blinks}] = noStones(cache, stone*2024, blinks-1)
	}
	return cache[xy{stone, blinks}]
}
