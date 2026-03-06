package nums

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/nbcx/kit/types"
)

// RandInt RandInt
func RandInt(min, max int) int { // nolint
	return min + rand.Intn(max-min+1)
}

// RandInt32 RandInt32
func RandInt32(min, max int32) int32 { // nolint
	return min + rand.Int31n(max-min+1)
}

// RandInt64 RandInt64
func RandInt64(min, max int64) int64 { // nolint
	return min + rand.Int63n(max-min+1)
}

// RandFloat 随机一个浮点数, place 指定保留的小数位数
func RandFloat[T float32 | float64](min, max T, place int) T { // nolint
	base := math.Pow10(place)
	return T(RandInt64(int64(float64(min)*base), int64(float64(max)*base))) / T(base)
}

// Rand 在指定区间里随机一个数, 结果可能包含min或max
func Rand[T types.Number](min, max T) T { // nolint
	l := 0
	if sp := strings.Split(fmt.Sprintf("%v", min), "."); len(sp) > 1 {
		l = len(sp[1])
	}
	if sp := strings.Split(fmt.Sprintf("%v", max), "."); len(sp) > 1 {
		if len(sp[1]) > l {
			l = len(sp[1])
		}
	}
	enlarge := T(math.Pow(10, float64(l)))
	minInt64, maxInt64 := int64(min*enlarge), int64(max*enlarge)

	return T(minInt64+rand.Int63n(maxInt64-minInt64+1)) / enlarge
}
