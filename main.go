package main

import (
	"fmt"

	"github.com/tobby-s/advent-of-code-2024/puzzles"
)

var sols = [][2]func() int{
	{puzzles.D1P1, puzzles.D1P2},
	{puzzles.D2P1, puzzles.D2P2},
	{puzzles.D3P1, puzzles.D3P2},
	{puzzles.D4P1, puzzles.D4P2},
	{puzzles.D5P1, puzzles.D5P2},
	{puzzles.D6P1, nil}, // d6p2 slow
	{puzzles.D7P1, puzzles.D7P2},
	{puzzles.D8P1, puzzles.D8P2},
	{puzzles.D9P1, puzzles.D9P2},
	{puzzles.D10P1, puzzles.D10P2},
	{puzzles.D11P1, puzzles.D11P2},
	{puzzles.D12P1, puzzles.D12P2},
}

func main() {
	for i, funcs := range sols {
		if funcs[0] != nil {
			fmt.Printf("Day %d Puzzle 1 Solution: %d\n", i+1, funcs[0]())
		}
		if funcs[1] != nil {
			fmt.Printf("Day %d Puzzle 2 Solution: %d\n", i+1, funcs[1]())
		}
	}
}
