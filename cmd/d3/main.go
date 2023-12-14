package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2023/fetch"
)

func main() {
	d := 3
	input := fetch.Day(d)

	fmt.Printf("# Day %d #\nP1: %s\nP2: %s\n", d, p1(input), p2(input))
}

type part struct {
	number int
	symbol string
}

type pos struct {
	x, y int
}

func p1(input string) string {
	input = strings.TrimSpace(input)
	symbols := map[pos]string{}

	for row, l := range strings.Split(input, "\n") {
		for col, c := range strings.Split(l, "") {
			if !(c >= "0" && c <= "9" || c == ".") {
				symbols[pos{row, col}] = c
			}
		}
	}

	parts := []part{}
	for row, l := range strings.Split(input, "\n") {
		ci := 0
		ds := ""
		neighborSymbol := ""
		for {
			isNumber := l[ci] >= '0' && l[ci] <= '9'

			if isNumber {
				ds += l[ci : ci+1]
				neighborSymbol += getNeighbourSymber(symbols, pos{row, ci})
			}

			if ds != "" &&
				(!isNumber ||
					ci == len(l)-1) && len(neighborSymbol) > 0 {
				number, err := strconv.Atoi(ds)
				if err != nil {
					panic(err)
				}

				parts = append(parts, part{
					number: number,
					symbol: neighborSymbol[:1],
				})
				ds = ""
				neighborSymbol = ""
			}

			if ci == len(l)-1 {
				break
			}
			ci++
		}
	}

	sum := 0
	for _, p := range parts {
		sum += p.number
		fmt.Printf("%s, %d\n", p.symbol, p.number)
	}

	return fmt.Sprintf("%d", sum)
}

func getNeighbourSymber(symbols map[pos]string, p pos) string {
	return symbols[pos{p.x - 1, p.y - 1}] +
		symbols[pos{p.x, p.y - 1}] +
		symbols[pos{p.x + 1, p.y - 1}] +
		symbols[pos{p.x - 1, p.y}] +
		symbols[pos{p.x + 1, p.y}] +
		symbols[pos{p.x - 1, p.y + 1}] +
		symbols[pos{p.x, p.y + 1}] +
		symbols[pos{p.x + 1, p.y + 1}]

}

func p2(input string) string {
	return "-1"
}
