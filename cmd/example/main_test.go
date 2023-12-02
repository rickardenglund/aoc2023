package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_p1(t *testing.T) {
	input := "hejho"
	require.Equal(t, "-1", p1(input))
}

func Test_p2(t *testing.T) {
	input := "hejhopp2"
	require.Equal(t, "-1", p2(input))
}
