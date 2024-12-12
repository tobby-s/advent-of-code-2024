package puzzles

import (
	"github.com/tobby-s/advent-of-code-2024/utils"
)

func D12P1() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/12/input")
	rows := len(data) - 1
	cols := len(data[0])

	regions := map[xy]int{}
	areas := map[int]int{}
	perims := map[int]int{}
	nextregion := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, ok := regions[xy{i, j}]; !ok {
				nextregion += 1
				plant := read(data, i, j)
				queue := []xy{{i, j}}
				for len(queue) > 0 {
					nextqueue := []xy{}
					for _, coord := range queue {
						if _, f := regions[coord]; !f {
							regions[coord] = nextregion
							areas[nextregion] += 1
							perims[nextregion] += 4
							for _, neighbouring := range []xy{{coord.x - 1, coord.y}, {coord.x + 1, coord.y}, {coord.x, coord.y - 1}, {coord.x, coord.y + 1}} {
								r, ok := regions[neighbouring]
								if !ok && read(data, neighbouring.x, neighbouring.y) == plant {
									nextqueue = append(nextqueue, neighbouring)
								} else if ok && r == nextregion {
									perims[nextregion] -= 2
								}
							}
						}
					}
					queue = nextqueue
				}
			}
		}
	}
	for k := range areas {
		result += areas[k] * perims[k]
	}
	return
}

func D12P2() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/12/input")
	rows := len(data) - 1
	cols := len(data[0])

	regions := map[xy]int{}
	areas := map[int]int{}
	corners := map[int]map[xy]int{}
	nextregion := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, ok := regions[xy{i, j}]; !ok {
				nextregion += 1
				plant := read(data, i, j)
				queue := []xy{{i, j}}
				for len(queue) > 0 {
					nextqueue := []xy{}
					for _, coord := range queue {
						if _, f := regions[coord]; !f {
							regions[coord] = nextregion
							areas[nextregion] += 1
							if _, ok := corners[nextregion]; !ok {
								corners[nextregion] = map[xy]int{}
							}
							corners[nextregion][coord] += 1
							corners[nextregion][xy{coord.x + 1, coord.y}] += 2
							corners[nextregion][xy{coord.x, coord.y + 1}] += 4
							corners[nextregion][xy{coord.x + 1, coord.y + 1}] += 8
							for _, neighbouring := range []xy{{coord.x - 1, coord.y}, {coord.x + 1, coord.y}, {coord.x, coord.y - 1}, {coord.x, coord.y + 1}} {
								_, ok := regions[neighbouring]
								if !ok && read(data, neighbouring.x, neighbouring.y) == plant {
									nextqueue = append(nextqueue, neighbouring)
								}
							}
						}
					}
					queue = nextqueue
				}
			}
		}
	}
	for k := range areas {
		c := corners[k]
		sides := 0
		for _, hash := range c {
			switch hash {
			case 6, 9:
				sides += 2
			case 14, 13, 11, 7, 1, 2, 4, 8:
				sides += 1
			}
		}
		result += areas[k] * sides
	}
	return
}
