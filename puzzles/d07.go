package puzzles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tobby-s/advent-of-code-2024/utils"

)

type equation struct {
	target   int
	summands []int
}

func parseData(data []string) []equation {
	eqns := make([]equation, 0)
	for _, row := range data {
		eqn := equation{}
		firstSep := strings.SplitN(row, ": ", 2)
		if len(firstSep) > 1 {
			target, _ := strconv.Atoi(firstSep[0])
			eqn.target = target
			nums := strings.Split(firstSep[1], " ")
			for _, numString := range nums {
				num, _ := strconv.Atoi(numString)
				eqn.summands = append(eqn.summands, num)
			}
			eqns = append(eqns, eqn)
		}
		}
	return eqns
}

func D7P1() (sum int) {
	data := utils.LoadData(`https://adventofcode.com/2024/day/7/input`)
	eqns := parseData(data)
	for _, eqn := range eqns {
		if canHitTarget(eqn.target, eqn.summands[0], eqn.summands[1:]) {
			sum += eqn.target
		}
	}
	return
}

func D7P2() (sum int) {
	data := utils.LoadData(`https://adventofcode.com/2024/day/7/input`)
	eqns := parseData(data)
	for _, eqn := range eqns {
		if canHitTarget2(eqn.target, eqn.summands[0], eqn.summands[1:]) {
			sum += eqn.target
		}
	}
	return
}

func canHitTarget(target, accumulator int, remaining []int) bool {
	if accumulator > target {
		return false
	}
	switch len(remaining) {
	case 0:
		return target == accumulator
	case 1:
		return (accumulator * remaining[0]) == target || (accumulator + remaining[0]) == target
	default:
		return canHitTarget(target, accumulator * remaining[0], remaining[1:]) || canHitTarget(target, accumulator + remaining[0], remaining[1:])
	}
}

func canHitTarget2(target, accumulator int, remaining []int) bool {
	if accumulator > target {
		return false
	}
	switch len(remaining) {
	case 0:
		return target == accumulator
	case 1:
		return (accumulator * remaining[0]) == target || (accumulator + remaining[0]) == target || concatenate(accumulator, remaining[0]) == target
	default:
		return canHitTarget2(target, accumulator * remaining[0], remaining[1:]) || canHitTarget2(target, accumulator + remaining[0], remaining[1:]) || canHitTarget2(target, concatenate(accumulator, remaining[0]), remaining[1:])
	}
}

func concatenate(a int, b int) int {
	str := fmt.Sprintf("%d%d", a, b)
	num , _ := strconv.Atoi(str)
	return num
}