package s

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/nbcx/kit/nums"
	"github.com/nbcx/kit/types"
)

// Weighted 根据权重随机，返回对应选项的索引，O(n)
func Weighted[T types.Integer](weightArray []T) (int, error) {
	total := Sum(weightArray)
	if total == 0 {
		return 0, fmt.Errorf("weight array total is zero in %+v", weightArray)
	}
	rv := rand.Int63n(int64(total))
	for i, v := range weightArray {
		v64 := int64(v)
		if rv < v64 {
			return i, nil
		}
		rv -= v64
	}
	return len(weightArray) - 1, nil
}

// WeightedChoiceFloat32 根据权重随机，返回对应选项的索引，O(n)
func WeightedChoiceFloat32(weights []float32) (int, error) {
	total := Sum(weights)
	if total == 0 {
		return 0, fmt.Errorf("weight array total is zero in %+v", weights)
	}
	if total < 0 || float32(math.Abs(float64(total-0.0))) < 1e-6 {
		panic("invalid argument to WeightedChoiceFloat32")
	}

	total *= rand.Float32()
	for i, v := range weights {
		if total < v {
			return i, nil
		}
		total -= v
	}
	return len(weights) - 1, nil
}

// WeightedChoiceFloat32 根据权重随机，返回对应选项的索引，O(n)
func WeightedChoiceFloat64(weights []float64) (int, error) {
	total := Sum(weights)
	if total == 0 {
		return 0, fmt.Errorf("weight array total is zero in %+v", weights)
	}
	if total < 0 || float32(math.Abs(float64(total-0.0))) < 1e-6 {
		panic("invalid argument to WeightedChoiceFloat32")
	}

	total *= rand.Float64()
	for i, v := range weights {
		if total < v {
			return i, nil
		}
		total -= v
	}
	return len(weights) - 1, nil
}

// RandChoice 随机从切片中选出n个元素
func RandChoice[T any](array []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if n == 1 {
		return []T{array[rand.Intn(len(array))]}
	}
	tmp := make([]T, len(array))
	copy(tmp, array)
	if len(tmp) <= n {
		return tmp
	}
	Shuffle(tmp)
	return tmp[:n]
}

// Combinations 从数组中选出m个任意组合
// 算法：先固定某一位的数字，再遍历其他位的可能性，递归此过程
func Combinations[T types.Integer](arr []T, m T) [][]T {
	length := T(len(arr))
	if arr == nil || m > length || m <= 0 {
		return nil
	}
	result := make([][]T, 0, C(length, m))
	data := make([]T, m)
	var helper func(T, T, T)

	helper = func(start T, end T, index T) {
		if index == m {
			d := make([]T, m)
			copy(d, data)
			result = append(result, d)
			return
		}
		for i := start; i < end && end-i+1 >= m-index; i++ {
			data[index] = arr[i]
			helper(i+1, end, index+1)
			// 去重
			for i+1 < end && arr[i] == arr[i+1] {
				i++
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	helper(0, length, 0)
	return result
}

// Perm 从给定的上限值total，生成指定n的随机数
func Perm[T types.Integer](total, n int) (result []T, err error) {
	if n > total || n == 0 {
		return nil, fmt.Errorf("perm total %v or n %v err", total, n)
	}
	// 确定生成算法
	if float64(n)/float64(total) > 0.5 {
		m := make([]T, total)
		for i := 0; i < total; i++ {
			j := rand.Intn(i + 1)
			m[i] = m[j]
			m[j] = T(i)
		}
		return m[:n], nil
	}
	resultM := make(map[T]struct{}, n)
	result = make([]T, n)
	for i := 0; i < n; {
		value := T(rand.Intn(total))
		if _, ok := resultM[value]; ok {
			continue
		}
		resultM[value] = struct{}{}
		result[i] = value
		i++
	}
	return
}

// C 计算组合结果
func C[T types.Integer](n, k T) T {
	i := k + 1
	r := n - k
	if r > k {
		i = r + 1
		r = k
	}
	f1, f2 := T(1), T(1)
	j := T(1)
	for ; i <= n; i++ {
		f1 *= i
		for ; j <= r; j++ {
			f2 *= j
			if f2 > f1 {
				j++
				break
			}
			if gcd := nums.Gcd(f1, f2); gcd > 1 {
				f1, f2 = nums.Div(f1, f2)
			}
		}
	}
	return f1 / f2
}
