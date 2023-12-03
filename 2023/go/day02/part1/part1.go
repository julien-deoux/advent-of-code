package part1

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
		if isValidGame(g) {
			sum += uint(g.Id)
		}
	}
	println(sum)
}

func isValidSet(set game.Set) bool {
	maxRed := uint(12)
	maxGreen := uint(13)
	maxBlue := uint(14)

	return set.Red <= maxRed && set.Green <= maxGreen && set.Blue <= maxBlue
}

func isValidGame(g game.Game) bool {
	for _, set := range g.Sets {
		if !isValidSet(set) {
			return false
		}
	}
	return true
}
