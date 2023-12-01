package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2023/fetch"
)

func main() {
	d := 1
	input := fetch.Day(d)

	headline := fmt.Sprintf("# Day %d #", d)
	fmt.Printf("%s\nP1: %s\nP2: %s\n", headline, p1(input), p2(input))
}

func p1(input string) string {
	return sumCalibrationValues(input, getDigit)
}

func p2(input string) string {
	combinedGetDigit := func(l string) *int {
		d := getDigit(l)
		if d != nil {
			return d
		}

		return getDigitString(l)
	}

	return sumCalibrationValues(input, combinedGetDigit)
}

func sumCalibrationValues(calibrationsDocument string, getDigitFuncs func(string) *int) string {
	sum := 0
	for _, l := range strings.Split(calibrationsDocument, "\n") {
		if l == "" {
			continue
		}

		digits := getDigits(l, getDigitFuncs)

		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return fmt.Sprintf("%d", sum)
}

func getDigits(l string, getDigitFromStart func(string) *int) []int {
	i := 0
	digits := []int{}
	for i < len(l) {
		d := getDigitFromStart(l[i:])
		if d != nil {
			digits = append(digits, *d)
		}

		i++
	}

	return digits
}

func getDigit(l string) *int {
	d, err := strconv.Atoi(l[:1])
	if err != nil {
		return nil
	}

	return &d
}

func getDigitString(l string) *int {
	for s, n := range numbers {
		if strings.HasPrefix(l, s) {
			return &n
		}
	}

	return nil
}

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
