package puzzles

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/tobby-s/advent-of-code-2024/utils"
)

var re = regexp.MustCompile(`(-?\d+)`)

func D14P1() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/14/input")
	var q1, q2, q3, q4 int
	for _, d := range data {
		if len(d) > 0 {
			matches := re.FindAllStringSubmatch(d, -1)
			p := xy{mustToInt(matches[0][0]), mustToInt(matches[1][0])}
			v := xy{mustToInt(matches[2][0]), mustToInt(matches[3][0])}
			x := (((p.x + v.x*100) % 101) + 101) % 101
			y := (((p.y + v.y*100) % 103) + 103) % 103
			switch {
			case x < 50 && y < 51:
				q2 += 1
			case x < 50 && y > 51:
				q3 += 1
			case x > 50 && y > 51:
				q4 += 1
			case x > 50 && y < 51:
				q1 += 1
			}
		}
	}
	return q1 * q2 * q3 * q4
}

func D14P2() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/14/input")
	robots := make([]*robot, len(data)-1)
	for i, l := range data {
		matches := re.FindAllStringSubmatch(l, -1)
		if len(matches) >= 4 {
			robots[i] = &robot{mustToInt(matches[0][0]), mustToInt(matches[1][0]), mustToInt(matches[2][0]), mustToInt(matches[3][0])}
		}
	}

	for i := 1; i < 10000; i++ {
		m := initMatrix(101, 103, " ")
		for _, r := range robots {
			r.step()
		}
		draw(robots, m)
		if hasStraightLine(m) {
			drawPrint(m)
			return i
		}
	}
	return
}

func mustToInt(s string) int {
	k, _ := strconv.Atoi(s)
	return k
}

type robot struct {
	x, y, px, py int
}

func (r *robot) step() {
	newx := (r.px + r.x) % 101
	if newx < 0 {
		newx += 101
	}
	newy := (r.y + r.py) % 103
	if newy < 0 {
		newy += 103
	}
	r.x = newx
	r.y = newy
}

func draw(robots []*robot, data [][]string) {
	for _, r := range robots {
		data[r.y][r.x] = "*"
	}
}

func drawPrint(data [][]string) {
	for _, line := range data {
		fmt.Println(strings.Join(line, ""))
	}
}

func hasStraightLine(data [][]string) bool {
	for _, line := range data {
		if len(line) > 0 {
			if strings.Contains(strings.Join(line, ""), "*********") {
				return true
			}
		}
	}
	return false
}

func initMatrix[T any](w, h int, def T) [][]T {
	m := make([][]T, h)
	for i := range m {
		m[i] = make([]T, w)
		for j := range m[i] {
			m[i][j] = def
		}
	}
	return m
}
