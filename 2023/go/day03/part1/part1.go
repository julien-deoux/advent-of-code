package part1

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	runFile("example.txt")
	runFile("input.txt")
}

func runFile(path string) {
	data, _ := os.ReadFile(path)
	lines := strings.Split(string(data), "\n")
	parts := make([]part, 0)
	numbers := make([]number, 0)

	for y, line := range lines {
		newParts, newNumbers := parseLine(line, y)
		parts = append(parts, newParts...)
		numbers = append(numbers, newNumbers...)
	}

	sum := int(0)
	for _, number := range numbers {
		if touchesPart(number, parts) {
			sum += number.Value
		}
	}

	println(sum)
}

func parseLine(line string, y int) ([]part, []number) {
	parts := make([]part, 0)
	numbers := make([]number, 0)
	currentNumber := ""
	for x, c := range line {
		if unicode.IsDigit(c) {
			currentNumber += string(c)
		} else {
			if currentNumber != "" {
				numbers = append(numbers, parseNumber(currentNumber, x, y))
				currentNumber = ""
			}
			if c != '.' {
				parts = append(parts, part{string(c), x, y})
			}
		}
	}
	if currentNumber != "" {
		numbers = append(numbers, parseNumber(currentNumber, len(line), y))
	}
	return parts, numbers
}

func parseNumber(numberStr string, x int, y int) number {
	xHigh := x - 1
	xLow := x - len(numberStr)
	value, _ := strconv.ParseInt(numberStr, 10, 64)
	return number{
		int(value),
		xLow,
		xHigh,
		y,
	}
}

func touchesPart(number number, parts []part) bool {
	for _, part := range parts {
		verticallyClose := number.Y-1 <= part.Y && part.Y <= number.Y+1
		horizontallyClose := number.XLow-1 <= part.X && part.X <= number.XHigh+1
		if verticallyClose && horizontallyClose {
			return true
		}
	}
	return false
}

type part struct {
	Symbol string
	X      int
	Y      int
}

type number struct {
	Value int
	XLow  int
	XHigh int
	Y     int
}
