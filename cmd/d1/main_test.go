package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n"

func Test_p1(t *testing.T) {
	require.Equal(t, p1(input), "142")
}

func Test_p2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\n"
	require.Equal(t, p2(input), "281")

}
