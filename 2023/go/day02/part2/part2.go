package part2

import (
	game "julien-deoux/day02/util"
	"os"
	"strings"
)

func Run() {
	runFile("example.txt")
	runFile("input.txt")
}

func runFile(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")
	sum := uint(0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		g := game.ParseGame(line)
		minimumSet := minimumNeededDie(g)
		sum += power(minimumSet)
	}
	println(sum)
}

func minimumNeededDie(g game.Game) game.Set {
	result := game.Set{}

	for _, set := range g.Sets {
		if set.Red > result.Red {
			result.Red = set.Red
		}
		if set.Green > result.Green {
			result.Green = set.Green
		}
		if set.Blue > result.Blue {
			result.Blue = set.Blue
		}
	}

	return result
}

func power(set game.Set) uint {
	return set.Red * set.Blue * set.Green
}
