package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2023/fetch"
)

func main() {
	d := 2
	input := fetch.Day(d)

	cs := CubeSet{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	fmt.Printf("# Day %d #\nP1: %s\nP2: %s\n", d, p1(input, cs), p2(input))
}

func p1(input string, cs CubeSet) string {
	games := getGames(input)

	sum := 0
	for _, g := range games {
		if g.possibleWith(cs) {
			sum += g.id
		}
	}

	return fmt.Sprintf("%d", sum)
}
func p2(input string) string {
	return ""
}

type CubeSet map[string]int

func (s CubeSet) count() int {
	total := 0
	for _, n := range s {
		total += n
	}

	return total
}

type game struct {
	id     int
	rounds []CubeSet
}

func (g game) possibleWith(all CubeSet) bool {
	for _, r := range g.rounds {
		if r.count() > all.count() {
			return false
		}

		for c, n := range r {
			if n > all[c] {
				return false

			}
		}
	}

	return true
}

func getGames(input string) []game {
	games := []game{}
	id := 1
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		i := strings.Index(l, ":")
		l = l[i+2:]

		rounds := []CubeSet{}
		for _, rs := range strings.Split(l, "; ") {
			cs := CubeSet{}
			for _, clrs := range strings.Split(rs, ", ") {
				clr := strings.Split(clrs, " ")
				n, err := strconv.Atoi(clr[0])
				if err != nil {
					panic(err)
				}

				cs[clr[1]] = n
			}

			rounds = append(rounds, cs)
		}

		games = append(games, game{id: id, rounds: rounds})
		id++
	}

	return games
}
