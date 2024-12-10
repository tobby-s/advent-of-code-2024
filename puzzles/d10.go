package puzzles

import (
	"github.com/tobby-s/advent-of-code-2024/utils"
)

type xy struct {
	x, y int
}

func D10P1() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/10/input")
	data = data[:(len(data) - 1)]
	rows := len(data)
	cols := len(data[0])
	scores := make([][]int, rows)
	for i := range scores {
		scores[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if data[i][j] == "0"[0] {
				result += trailHeadScore(data, i, j, scores)
			}
		}
	}
	return
}

func trailHeadScore(data []string, i, j int, _ [][]int) (result int) {
	current := read(data, i, j)
	queue := []xy{{i, j}}
	for current < "9"[0] {
		nextqueue := []xy{}
		for _, k := range queue {
			for _, l := range []xy{
				{k.x - 1, k.y},
				{k.x + 1, k.y},
				{k.x, k.y - 1},
				{k.x, k.y + 1},
			} {
				if read(data, l.x, l.y) == current+1 {
					nextqueue = append(nextqueue, l)
				}
			}
		}
		if current == "8"[0] {
			m := make(map[xy]bool)
			for _, n := range nextqueue {
				m[n] = true
			}
			return len(m)
		} else {
			current += 1
			queue = nextqueue
		}
	}
	return
}

func trailHeadRating(data []string, x, y int, scores [][]int) (result int) {
	if scores[x][y] > 0 {
		return scores[x][y] - 1
	}
	currentSpot := data[x][y]
	if currentSpot == "9"[0] {
		scores[x][y] = 2
		return 1
	}
	up := read(data, x-1, y)
	if up == currentSpot+1 {
		result += trailHeadRating(data, x-1, y, scores)
	}
	down := read(data, x+1, y)
	if down == currentSpot+1 {
		result += trailHeadRating(data, x+1, y, scores)
	}
	left := read(data, x, y-1)
	if left == currentSpot+1 {
		result += trailHeadRating(data, x, y-1, scores)
	}
	right := read(data, x, y+1)
	if right == currentSpot+1 {
		result += trailHeadRating(data, x, y+1, scores)
	}
	scores[x][y] = result + 1
	return
}

func D10P2() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/10/input")
	data = data[:(len(data) - 1)]
	rows := len(data)
	cols := len(data[0])
	scores := make([][]int, rows)
	for i := range scores {
		scores[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if data[i][j] == "0"[0] {
				result += trailHeadRating(data, i, j, scores)
			}
		}
	}
	return
}
