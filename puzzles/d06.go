package puzzles

import (
	"fmt"

	"github.com/tobby-s/advent-of-code-2024/utils"

)

var dirs = [4][2]int {{-1,0}, {0,1}, {1,0}, {0,-1}}

func D6P1() (result int) {
	envMap := utils.LoadData("https://adventofcode.com/2024/day/6/input")
	visited := map[string]bool{}
	dirs := [4][2]int {{-1,0}, {0,1}, {1,0}, {0,-1}}
	dir := 0
	x,y := findGuard(envMap)
	defer func() {
		if r := recover(); r != nil {
			result = len(visited)
		}
	}()
	for true {
		visited[fmt.Sprintf("%03d%03d",x,y)] = true
		momentum := dirs[dir]
		if envMap[x + momentum[0]][y+momentum[1]] == "#"[0] {
			dir = (dir + 1) % 4
			continue
		} else {
			x += momentum[0]
			y += momentum[1]
			continue
		}
	}
	return
}

func D6P2() (result int) {
	envMap := utils.LoadData("https://adventofcode.com/2024/day/6/input")
	x,y := findGuard(envMap)
	obstacles := map[string]bool{fmt.Sprintf("%03d%03d",x,y):false}
	dir := 0
	defer func() {
		if r := recover(); r != nil {
			for _, v := range obstacles {
				if v {
					result += 1
				}
			}
		}
	}()

	for true{
		momentum := dirs[dir]
		// check whether already at end of map first
		x2, y2, dir2 := step(envMap,x,y,dir)
		nextPosKey := fmt.Sprintf("%03d%03d",x + momentum[0],y + momentum[1])
		if _, ok := obstacles[nextPosKey]; !ok {
			if willTurningHereCauseInfLoop(envMap, x, y , dir) {
				obstacles[nextPosKey] = true
			} else {
				obstacles[nextPosKey] = false
			}
		}
		x, y, dir = x2, y2, dir2
	}
	return
}

func findGuard(envMap []string) (int, int) {
	for x, row := range envMap {
		for y, b := range []byte(row) {
			if b == "^"[0] {
				return x, y
			}
		}
	}
	panic("didn't find guard")
}

func step(envMap []string, x, y, dir int) (int, int, int) {
	momentum := dirs[dir]
	if envMap[x + momentum[0]][y+momentum[1]] == "#"[0] {
		return x, y, (dir + 1) % 4
	}
	return x + momentum[0], y + momentum[1], dir
}

func stepWithAddedObstacle(envMap []string, x, y, dir, obsx, obsy int)(int, int, int) {
	momentum := dirs[dir]
	if envMap[x + momentum[0]][y+momentum[1]] == "#"[0] || (x + momentum[0] == obsx && y + momentum[1] == obsy) {
		return x, y, (dir + 1) % 4
	}
	return x + momentum[0], y + momentum[1], dir
}

func willTurningHereCauseInfLoop (envMap []string, x, y, dir int) bool {
	visited := map[string]bool{fmt.Sprintf("%03d%03d%d",x,y,dir): true}
	obsx, obsy := x + dirs[dir][0], y + dirs[dir][1]
	dir = (dir + 1) % 4
	defer func() {
	recover()
	}()
	for true {
		visited[fmt.Sprintf("%03d%03d%d",x,y,dir)] = true
		x, y, dir = stepWithAddedObstacle(envMap, x, y, dir, obsx, obsy)
		if _, ok := visited[fmt.Sprintf("%03d%03d%d",x,y,dir)]; ok {
			return true
		}
	}
	return false
}