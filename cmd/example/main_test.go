package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_p1(t *testing.T) {
	input := "hejho"
	require.Equal(t, p1(input), "-1")
}

func Test_p2(t *testing.T) {
	input := "hejhopp2"
	require.Equal(t, p2(input), "-1")
}
