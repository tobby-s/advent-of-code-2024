package puzzles

import (
	"reflect"

	"github.com/tobby-s/advent-of-code-2024/utils"
)

func D4P1() int {
	url := `https://adventofcode.com/2024/day/4/input`
	data := utils.LoadData(url)
	dirs := [][2]int{{-1,1},{-1,0},{-1,-1},{0,-1},{0,1}, {1,0}, {1,-1}, {1,1}}
	sum := 0
	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[x]); y++ {
			if read(data, x, y) == "X"[0] {
				for _, dir := range dirs {
					if checkDir(data, x, y, dir) {
						sum += 1
					}
				}
			}
		}
	}
	return sum
}

func D4P2() int {
	url := `https://adventofcode.com/2024/day/4/input`
	data := utils.LoadData(url)
	sum := 0
	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[x]); y++ {
			if checkXmas(data, x, y) {
				sum += 1
			}
		}
	}
	return sum
}

func read(data []string, x, y int) byte {
	var b byte
	defer func() {
		recover()
	}()
	b = data[x][y]
	return b
}

func checkDir(data []string, startx, starty int, dir [2]int) bool {
	for i, b := range []byte("MAS") {
		if read(data, startx + (i+1) * dir[0], starty + (i+1) * dir[1]) != b {
			return false
		}
	}
	return true
}
func checkXmas(data []string, startx, starty int) bool {
	if read(data, startx, starty) != "A"[0] {
		return false
	}
	diagLeft := map[byte]int{}
	diagRight := map[byte]int{}
	want := map[byte]int{"M"[0]: 1, "S"[0]: 1}
	for _, dir := range [2][2]int{{1,1}, {-1,-1}} {
		diagRight[read(data, startx + dir[0], starty + dir[1])] += 1
	}
	for _, dir := range [2][2]int{{1,-1}, {-1,1}} {
		diagLeft[read(data, startx + dir[0], starty + dir[1])] += 1
	}
		return reflect.DeepEqual(diagLeft, want) && reflect.DeepEqual(diagRight, want)
}