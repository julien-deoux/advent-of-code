package part2

import (
	"os"
	"strconv"
	"strings"
)

func Run() {
	runFile("part2/example.txt")
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
	valueString := findFirstAndLastDigits(line)
	value, err := strconv.ParseUint(valueString, 10, 64)
	if err != nil {
		return 0
	}
	return uint(value)
}

func findFirstAndLastDigits(s string) string {
	crawler1 := crawler{"1", 0}
	crawler2 := crawler{"2", 0}
	crawler3 := crawler{"3", 0}
	crawler4 := crawler{"4", 0}
	crawler5 := crawler{"5", 0}
	crawler6 := crawler{"6", 0}
	crawler7 := crawler{"7", 0}
	crawler8 := crawler{"8", 0}
	crawler9 := crawler{"9", 0}
	crawlerOne := crawler{"one", 0}
	crawlerTwo := crawler{"two", 0}
	crawlerThree := crawler{"three", 0}
	crawlerFour := crawler{"four", 0}
	crawlerFive := crawler{"five", 0}
	crawlerSix := crawler{"six", 0}
	crawlerSeven := crawler{"seven", 0}
	crawlerEight := crawler{"eight", 0}
	crawlerNine := crawler{"nine", 0}

	var firstDigit byte
	var lastDigit byte

	checkDigit := func(cr *crawler, digit byte, result byte) {
		if cr.crawl(digit) {
			if firstDigit == 0 {
				firstDigit = result
			}
			lastDigit = result
		}
	}

	for i := 0; i < len(s); i++ {
		checkDigit(&crawler1, s[i], '1')
		checkDigit(&crawler2, s[i], '2')
		checkDigit(&crawler3, s[i], '3')
		checkDigit(&crawler4, s[i], '4')
		checkDigit(&crawler5, s[i], '5')
		checkDigit(&crawler6, s[i], '6')
		checkDigit(&crawler7, s[i], '7')
		checkDigit(&crawler8, s[i], '8')
		checkDigit(&crawler9, s[i], '9')
		checkDigit(&crawlerOne, s[i], '1')
		checkDigit(&crawlerTwo, s[i], '2')
		checkDigit(&crawlerThree, s[i], '3')
		checkDigit(&crawlerFour, s[i], '4')
		checkDigit(&crawlerFive, s[i], '5')
		checkDigit(&crawlerSix, s[i], '6')
		checkDigit(&crawlerSeven, s[i], '7')
		checkDigit(&crawlerEight, s[i], '8')
		checkDigit(&crawlerNine, s[i], '9')
	}

	return string([]byte{firstDigit, lastDigit})
}

type crawler struct {
	reference    string
	currentIndex int
}

func (c *crawler) crawl(chr byte) bool {
	if c.reference[c.currentIndex] != chr {
		if c.currentIndex == 0 {
			return false
		}
		c.currentIndex = 0
		return c.crawl(chr)
	}
	if c.currentIndex == len(c.reference)-1 {
		c.currentIndex = 0
		return true
	}
	c.currentIndex = c.currentIndex + 1
	return false
}
