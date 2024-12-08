package puzzles

import (
	"slices"
	"strconv"
	"strings"

	"github.com/tobby-s/advent-of-code-2024/utils"

)

func D5P1() int {
	data := utils.LoadData("https://adventofcode.com/2024/day/5/input")
	split := slices.Index(data, "")
	rules := data[:split]
	updates := data[(split+1):]

	ruleFuncs := make([]func(map[string]int)bool,len(rules))
	for i, rule := range rules {
		ruleFuncs[i] = parseRule(rule)
	}

	sum := 0
	for _, update := range updates {
		updateSlice := strings.Split(update, ",")
		if testPageOrder(updateSlice, ruleFuncs) {
			middlePage := len(updateSlice) / 2
			num, _ := strconv.Atoi(updateSlice[middlePage])
			sum += num
		}
	}

	return sum
}

func D5P2() int {
	data := utils.LoadData("https://adventofcode.com/2024/day/5/input")
	split := slices.Index(data, "")
	rules := data[:split]
	updates := data[(split+1):]

	ruleFuncs := make([]func(map[string]int)bool,len(rules))
	for i, rule := range rules {
		ruleFuncs[i] = parseRule(rule)
	}

	cmp := makeComparisonFunction(rules)

	sum := 0
	for _, update := range updates {
		updateSlice := strings.Split(update, ",")
		if !testPageOrder(updateSlice, ruleFuncs) {
			slices.SortFunc(updateSlice, cmp)
			middleIndex := len(updateSlice) / 2
			middlePage, _ := strconv.Atoi(updateSlice[middleIndex])
			sum += middlePage
		}
	}

	return sum
}

func makeComparisonFunction(rules []string) func(a, b string) int {
	ordering := map[struct{l, g string}]struct{}{}
	for _, rule := range rules {
		pages := strings.Split(rule, "|")
		ordering[struct{l, g string}{pages[0], pages[1]}] = struct{}{}
	}
	return func(a, b string) int {
	if _, ok := ordering[struct{l,g string}{a,b}]; ok {
		return -1
	} else if _, ok := ordering[struct{l,g string}{b,a}]; ok {
		return 1
	} else {
		return 0
	}
	}
}

func parseRule(rule string) func(map[string]int) bool {
	pages := strings.Split(rule, "|")
	return func(update map[string]int) bool {
		i1, ok1 := update[pages[0]]
		i2, ok2 := update[pages[1]]
		return !ok1 || !ok2 || i1 < i2 
	}
}

func testPageOrder(update []string, rules []func(map[string]int)bool) bool {
	order := map[string]int{}
	for i, p := range update {
		order[p] = i
	}
	for _, rule := range rules {
		if !rule(order) {
			return false
		}
	}
	return true
} 