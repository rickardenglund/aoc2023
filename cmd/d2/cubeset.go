package main

type CubeSet map[string]int

func (s CubeSet) count() int {
	total := 0
	for _, n := range s {
		total += n
	}

	return total
}

func (s CubeSet) Power() int {
	power := 1
	for _, n := range s {
		power *= n
	}

	return power
}
