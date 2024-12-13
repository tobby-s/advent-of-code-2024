package puzzles

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"

	"github.com/tobby-s/advent-of-code-2024/utils"
)

type machine struct {
	buttonA xy
	buttonB xy
	prize   xy
}

var parse = regexp.MustCompile(`(\d+)`)

func D13P1() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/13/input")
	machines := make([]machine, len(data)/4)
	for i := 0; i < len(data)/4; i++ {
		m := machine{}
		for j := 0; j < 4; j++ {
			nums := parse.FindAllStringSubmatch(data[i*4+j], -1)
			if len(nums) >= 2 {
				x, _ := strconv.Atoi(nums[0][0])
				y, _ := strconv.Atoi(nums[1][0])
				switch j {
				case 0:
					m.buttonA = xy{x, y}
				case 1:
					m.buttonB = xy{x, y}
				case 2:
					m.prize = xy{x, y}
				}
			}
		}
		machines[i] = m
	}
	for _, m := range machines {
		possibleCosts := []int{}
		aCount := 0
		for aCount*m.buttonA.x <= m.prize.x && aCount*m.buttonA.y <= m.prize.y {
			for bCount := 0; bCount*m.buttonB.x+m.buttonA.x*aCount <= m.prize.x &&
				bCount*m.buttonB.y+aCount*m.buttonA.y <= m.prize.y; bCount++ {
				if bCount*m.buttonB.x+m.buttonA.x*aCount == m.prize.x &&
					bCount*m.buttonB.y+aCount*m.buttonA.y == m.prize.y {
					possibleCosts = append(possibleCosts, aCount*3+bCount)
				}
			}
			aCount += 1
		}
		if len(possibleCosts) > 0 {
			fmt.Println(possibleCosts, m)
			result += slices.Min(possibleCosts)
		}
	}
	return
}

func D13P2() (result int) {
	data := utils.LoadData("https://adventofcode.com/2024/day/13/input")
	machines := make([]machine, len(data)/4)
	for i := 0; i < len(data)/4; i++ {
		m := machine{}
		for j := 0; j < 4; j++ {
			nums := parse.FindAllStringSubmatch(data[i*4+j], -1)
			if len(nums) >= 2 {
				x, _ := strconv.Atoi(nums[0][0])
				y, _ := strconv.Atoi(nums[1][0])
				switch j {
				case 0:
					m.buttonA = xy{x, y}
				case 1:
					m.buttonB = xy{x, y}
				case 2:
					m.prize = xy{x + 10000000000000, y + 10000000000000}
				}
			}
		}
		machines[i] = m
	}
	for _, m := range machines {
		det := m.buttonA.x*m.buttonB.y - m.buttonA.y*m.buttonB.x
		if det != 0 {
			a := float64(m.prize.x*m.buttonB.y-m.prize.y*m.buttonB.x) / float64(det)
			b := float64(m.prize.y*m.buttonA.x-m.prize.x*m.buttonA.y) / float64(det)
			if math.Floor(a) == a && math.Floor(b) == b && a >= 0 && b >= 0 {
				result += 3*int(a) + int(b)
				continue
			}
		} else {
			possibleCosts := []int{}
			if g := gcd(m.buttonA.x, m.buttonB.x); m.prize.x%g == 0 {
				for aCount := 0; aCount*m.buttonA.x <= m.prize.x; aCount++ {
					if m.prize.x-aCount*m.buttonA.x%m.buttonB.x == 0 {
						possibleCosts = append(possibleCosts, aCount*3+(m.prize.x-aCount*m.buttonA.x/m.buttonB.x))
						break
					}
				}
				for bCount := 0; bCount*m.buttonB.x <= m.prize.x; bCount++ {
					if m.prize.x-bCount*m.buttonB.x%m.buttonA.x == 0 {
						possibleCosts = append(possibleCosts, bCount+(m.prize.x-bCount*m.buttonB.x)/m.buttonA.x*3)
						break
					}
				}
			}
			if len(possibleCosts) > 0 {
				result += slices.Min(possibleCosts)
			}
		}
	}
	return
}
