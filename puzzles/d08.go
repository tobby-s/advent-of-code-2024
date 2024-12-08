package puzzles

import (
	"github.com/tobby-s/advent-of-code-2024/utils"

)

type coords struct {
	x, y int
}

func D8P1() int {
	data := utils.LoadData(`https://adventofcode.com/2024/day/8/input`)
	antennae := map[byte][]coords{}
	antinodes := map[coords]struct{}{}
	for x, row := range data {
		for y, b := range []byte(row) {
			if b != "."[0] {
				antennae[b] = append(antennae[b], coords{x, y})
			}
		}
	}
	for x, row := range data {
		for y, _ := range []byte(row) {
			for freq, ants := range antennae {
				if isAntinode(data, x, y, freq, ants) {
					antinodes[coords{x,y}] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}

func D8P2() int {
	data := utils.LoadData(`https://adventofcode.com/2024/day/8/input`)
	antennae := map[byte][]coords{}
	for x, row := range data {
		for y, b := range []byte(row) {
			if b != "."[0] {
				antennae[b] = append(antennae[b], coords{x, y})
			}
		}
	}
	antinodes := populateAntinodes(len(data) - 1, len(data[0]), antennae)
	return len(antinodes)
}

func isAntinode(data []string, x, y int, freq byte, ants []coords) (result bool){
	for _, antenna := range ants {
		if antenna.x != x || antenna.y != y {
			if read(data, antenna.x * 2 - x, antenna.y * 2 - y) == freq {
				result = true
			}
			if read(data, x * 3 - 2 * antenna.x, y * 3 - 2 * antenna.y) == freq {
				result = true
			}
		}
		if result {
			break
		}
	}
	return
}

func populateAntinodes(maxRows, maxCols int, antennae map[byte][]coords) map[coords]string {
	antinodes := map[coords]string{}
	for freq, ants := range antennae {
		if len(ants) < 2 {
			continue
		}
		for i, ant := range ants {
			for j := i+1; j < len(ants); j++ {
				diff := coords{ants[j].x - ant.x, ants[j].y - ant.y}
				// diff = reduce(diff)
				for k := 0;  k >= 0; k++ {
					scanCoords := coords{ant.x + k * diff.x, ant.y + k * diff.y}
					if scanCoords.x < 0 || scanCoords.y < 0 || scanCoords.x >= maxRows || scanCoords.y >= maxCols {
						break
					}
					antinodes[scanCoords] = string(freq)
				}
				for k := 0;  k >= 0; k++ {
					scanCoords := coords{ant.x - k * diff.x, ant.y - k * diff.y}
					if scanCoords.x < 0 || scanCoords.y < 0 || scanCoords.x >= maxRows || scanCoords.y >= maxCols {
						break
					}
					antinodes[scanCoords] = string(freq)
				}
			}
		}
	}
	return antinodes
}

func reduce(c coords) coords {
	g := gcd(c.x, c.y)
	return coords{c.x / g, c.y/g}
}

func gcd(a, b int) int {
	if a < 0 { a = -a}
	if b < 0 { b = -b}
	if b > a { a,b = b,a }
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}