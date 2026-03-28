package f

// If expr is true, return a; otherwise return b.
func If[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}

// Or expr is true, run f
func Or(expr bool, f func()) {
	if expr {
		f()
	}
}
