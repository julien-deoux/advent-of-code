package game

import (
	"strconv"
	"strings"
)

type Set struct {
	Red   uint
	Green uint
	Blue  uint
}

type Game struct {
	Id   uint
	Sets map[int]Set
}

func ParseGame(gameString string) Game {
	idAndSets := strings.Split(gameString, ": ")
	id, _ := strconv.ParseUint(strings.Split(idAndSets[0], " ")[1], 10, 64)
	setStrings := strings.Split(idAndSets[1], "; ")

	sets := make(map[int]Set)
	for i, setString := range setStrings {
		sets[i] = parseSet(setString)
	}

	return Game{uint(id), sets}
}

func parseSet(setString string) Set {
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

	return Set{red, green, blue}
}
