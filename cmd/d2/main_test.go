package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_p1(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green\n"
	cs := CubeSet{
		"red":   12,
		"green": 14,
		"blue":  14,
	}
	require.Equal(t, "8", p1(input, cs))
}

func Test_p2(t *testing.T) {
	input := "hejhopp2"
	require.Equal(t, "-1", p2(input))
}
