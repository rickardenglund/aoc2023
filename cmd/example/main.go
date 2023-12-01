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
	acc := 0
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		food, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}

		acc += food
	}

	return fmt.Sprintf("%d", acc)
}

func p2(input string) string {
	return fmt.Sprintf("%d", len(input))
}
