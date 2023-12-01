package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		handleLine(line)
	}
}

func handleLine(line string) {
	numbers := strings.Split(line, " ")
	var sum int64 = 0
	for _, number := range numbers {
		number, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return
		}
		sum += number
	}
	println(sum)
}
