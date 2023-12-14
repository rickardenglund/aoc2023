package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_p1(t *testing.T) {
	input := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..\n"
	require.Equal(t, "4361", p1(input))
}

func Test_p2(t *testing.T) {
	input := "hejhopp2"
	require.Equal(t, "-1", p2(input))
}
