package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		handleLine(line)
	}
}

func handleLine(line string) {
	numbers := strings.Split(line, " ")
	sum := int(0)
	for _, number := range numbers {
		number, _ := strconv.ParseInt(number, 10, 64)
		sum += int(number)
	}
	println(sum)
}
