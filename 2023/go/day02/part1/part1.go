package part1

import (
	"os"
	"strconv"
	"strings"
)

func Run() {
	runFile("part1/example.txt")
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
		game := parseGame(line)
		if game.isValid() {
			sum += uint(game.id)
		}
	}
	println(sum)
}

type set struct {
	red   uint
	green uint
	blue  uint
}

type game struct {
	id   uint
	sets map[int]set
}

func parseGame(gameString string) game {
	idAndSets := strings.Split(gameString, ": ")
	id, _ := strconv.ParseUint(strings.Split(idAndSets[0], " ")[1], 10, 64)
	setStrings := strings.Split(idAndSets[1], "; ")

	sets := make(map[int]set)
	for i, setString := range setStrings {
		sets[i] = parseSet(setString)
	}

	return game{uint(id), sets}
}

func parseSet(setString string) set {
	values := strings.Split(setString, ", ")
	var red, green, blue uint

	for _, value := range values {
		countAndColor := strings.Split(value, " ")
		count, _ := strconv.ParseUint(countAndColor[0], 10, 64)
		color := countAndColor[1]
		switch color {
		case "red":
			red = uint(count)
		case "green":
			green = uint(count)
		case "blue":
			blue = uint(count)
		}
	}

	return set{red, green, blue}
}

func (set set) isValid() bool {
	maxRed := uint(12)
	maxGreen := uint(13)
	maxBlue := uint(14)

	return set.red <= maxRed && set.green <= maxGreen && set.blue <= maxBlue
}

func (game game) isValid() bool {
	for _, set := range game.sets {
		if !set.isValid() {
			return false
		}
	}
	return true
}
