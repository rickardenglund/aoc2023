package main

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

func (g game) smallestStart() CubeSet {
	s := CubeSet{}
	for _, g := range g.rounds {
		for c, n := range g {
			if n > s[c] {
				s[c] = n
			}
		}
	}

	return s
}
