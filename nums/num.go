package nums

import (
	"math"

	"github.com/nbcx/kit/types"
)

// MinInt MinInt
func Min[T types.Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// MaxInt MaxInt
func Max[T types.Number](x, y T) T {
	if x < y {
		return y
	}
	return x
}

// Gcd Gcd
func Gcd[T types.Integer](a, b T) T {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// Div : divide by Gcd
func Div[T types.Integer](a, b T) (a0, b0 T) {
	gcd := Gcd(a, b)
	a /= gcd
	b /= gcd
	return a, b
}

// Range .
func Range[T types.Integer](start, end T) []T {
	if start >= end {
		return nil
	}
	result := make([]T, 0, end-start)
	for start < end {
		result = append(result, start)
		start++
	}
	return result
}

func IsEqualFloat32(l, r float32) bool {
	return float32(math.Abs(float64(l-r))) < 1e-6
}

func IsEqualFloat64(l, r float64) bool {
	return math.Abs(l-r) < 1e-15
}

// IsBetween 判断是否在区间内
func IsBetween[T types.Number](val T, between []T) bool {
	return val >= between[0] && val <= between[1]
}
