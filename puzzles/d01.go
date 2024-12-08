package puzzles

import (
	"slices"
	"strconv"
	"strings"

	"github.com/tobby-s/advent-of-code-2024/utils"
)
var url = "https://adventofcode.com/2024/day/1/input"

func D1P1() int {
	data := utils.LoadData(url)
	var list1, list2 []int
	for _, s := range data {
		ids := strings.Split(s, "   ")
		if len(ids) >= 2 {
			num1, _ := strconv.Atoi(ids[0])
			num2, _ := strconv.Atoi(ids[1])
			list1 = append(list1,num1)
			list2 = append(list2,num2)
			slices.Sort(list1)
			slices.Sort(list2)
		}
	}
	sum := 0
	for i, _ := range list1 {
		if list1[i] < list2[i] {
			sum = sum + list2[i] - list1[i]
		} else {
			sum = sum + list1[i] - list2[i]
		}
	}
	return sum
}

func D1P2() int {
	data := utils.LoadData(url)

	map1, map2 := map[int]int{}, map[int]int{}
	for _,s := range data {
		ids := strings.Split(s, "   ")
		if len(ids) >= 2 {
			num1, _ := strconv.Atoi(ids[0])
			num2, _ := strconv.Atoi(ids[1])
			map1[num1] += 1
			map2[num2] += 1
		}
	}
	simScore := 0
	for k, v := range map1 {
		simScore += k * v * map2[k]
	}
	return simScore
}