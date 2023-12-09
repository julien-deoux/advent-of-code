package part2

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
	parts := make([](*part), 0)
	numbers := make([]number, 0)

	for y, line := range lines {
		newParts, newNumbers := parseLine(line, y)
		parts = append(parts, newParts...)
		numbers = append(numbers, newNumbers...)
	}

	sum := int(0)
	for _, number := range numbers {
		registerToParts(number, parts)
	}

	for _, part := range parts {
		sum += gearRatio(*part)
	}

	println(sum)
}

func parseLine(line string, y int) ([](*part), []number) {
	parts := make([](*part), 0)
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
				parts = append(parts, &part{string(c), x, y, make([]number, 0)})
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

func registerToParts(number number, parts [](*part)) {
	for _, part := range parts {
		verticallyClose := number.Y-1 <= part.Y && part.Y <= number.Y+1
		horizontallyClose := number.XLow-1 <= part.X && part.X <= number.XHigh+1
		if verticallyClose && horizontallyClose {
			part.Numbers = append(part.Numbers, number)
		}
	}
}

func gearRatio(part part) int {
	if len(part.Numbers) != 2 {
		return 0
	}
	return part.Numbers[0].Value * part.Numbers[1].Value
}

type part struct {
	Symbol  string
	X       int
	Y       int
	Numbers []number
}

type number struct {
	Value int
	XLow  int
	XHigh int
	Y     int
}
