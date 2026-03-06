package nums

import (
	"math/rand"

	"github.com/nbcx/kit/types"
)

// HitRate100 是否命中百分比
func HitRate100(rate int) bool {
	return rand.Intn(100) < rate
}

// HitRate1000 是否命中千分比
func HitRate1000(rate int) bool {
	return rand.Intn(1000) < rate
}

// HitRate10000 是否命中万分比
func HitRate10000(rate int) bool {
	return rand.Intn(10000) < rate
}

// HitRate 是否命中概率
func HitRate[T types.Float](rate T) bool {
	return rand.Float64() < float64(rate)
}
