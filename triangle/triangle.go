package triangle

import "sort"

type Kind int

const (
	Equ Kind = iota
	Iso
	NaT
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	s := []float64{a, b, c}

	sort.Float64s(s)

	if s[0] <= 0 || // zero-length or negative side
		s[0] != s[0] || // NaN
		s[0]+s[1] <= s[2] { // triangle inequality
		return NaT
	}

	if s[0] == s[2] {
		return Equ
	}

	if s[0] == s[1] ||
		s[1] == s[2] {
		return Iso
	}

	return Sca
}
