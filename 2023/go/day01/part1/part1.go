package part1

import (
	"os"
	"regexp"
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
		calibrationValue := getCalibrationValue(line)
		sum += calibrationValue
	}
	println(sum)
}

func getCalibrationValue(line string) uint {
	digitRegex, err := regexp.Compile("[0-9]")
	if err != nil {
		panic(err)
	}

	digits := digitRegex.FindAllString(line, -1)
	if len(digits) == 0 {
		return 0
	}
	valueString := digits[0] + digits[len(digits)-1]
	value, err := strconv.ParseUint(valueString, 10, 64)
	if err != nil {
		return 0
	}
	return uint(value)
}
