package s

import (
	"math/rand"

	"github.com/nbcx/kit/types"
)

// Sum 求和
func Sum[T types.Number](data []T) T {
	var sum T
	for _, v := range data {
		sum += v
	}
	return sum
}

// RemoveValues 移除一些值
func RemoveValues[T comparable](data []T, values ...T) []T {
	toDelete := NewSet(values...)
	result := make([]T, 0, len(data))
	for _, v := range data {
		if !toDelete.Contains(v) {
			result = append(result, v)
		}
	}
	return result
}

// Remove 移除指定位置的值
func Remove[T any](data []T, index int) []T {
	if index < 0 || index >= len(data) {
		return data
	}
	return append(data[:index], data[index+1:]...)
}

// RemoveFirstValue 移除第一个指定的值
func RemoveFirstValue[T comparable](data []T, value T) []T {
	for i, d := range data {
		if d == value {
			return append(data[:i], data[i+1:]...)
		}
	}
	return data
}

// FindValueIndex 查找指定值的索引
func FindValueIndex[T comparable](data []T, value T) int {
	for i, d := range data {
		if d == value {
			return i
		}
	}
	return -1
}

// Contains 是否包含指定的值
func Contains[T comparable](data []T, value T) bool {
	for _, v := range data {
		if value == v {
			return true
		}
	}
	return false
}

// ContainsAll 是否包含所有指定的值
func ContainsAll[T comparable](data []T, values ...T) bool {
	for _, v := range values {
		if !Contains(data, v) {
			return false
		}
	}
	return true
}

// ContainsEqual 是否全等
func ContainsEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	// Use a map to count frequency of values in slice a
	freq := make(map[T]int)
	for _, val := range a {
		freq[val]++
	}
	// Subtract frequencies for slice b
	for _, val := range b {
		freq[val]--
		if freq[val] == 0 {
			delete(freq, val)
		}
	}
	// If map is empty, slices have same elements with same frequency
	return len(freq) == 0
}

// Reverse 翻转切片
func Reverse[T any](data []T) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// Shuffle 打乱数据
func Shuffle[T any](array []T) {
	for i := range array {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

// Pick 挑选出指定索引位置的值
func Pick[T any](data []T, indexes []int) []T {
	var result []T
	dataLen := len(data)
	for _, index := range indexes {
		if index < dataLen {
			result = append(result, data[index])
		}
	}
	return result
}

// Permutations 全排列
func Permutations(arr []int) [][]int {
	var helper func([]int, int)
	var res [][]int

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

// DirectProduct 任意多个集合的笛卡尔积（直积）
// 回溯法遍历所有可能性
func DirectProduct(items ...[]int) [][]int {
	if len(items) == 0 {
		return nil
	}
	size := 1
	for _, item := range items {
		size *= len(item)
	}
	result := make([][]int, 0, size)
	data := make([]int, len(items))
	var backtrack func(int)
	backtrack = func(index int) {
		if len(items) == index {
			d := make([]int, len(items))
			copy(d, data)
			result = append(result, d)
			return
		}
		for i := 0; i < len(items[index]); i++ {
			data[index] = items[index][i]
			backtrack(index + 1)
		}
	}
	backtrack(0)
	return result
}

// DirectProductInt32 任意多个集合的笛卡尔积（直积），Int32版
// 回溯法遍历所有可能性
func DirectProductInt32(items ...[]int32) [][]int32 {
	if len(items) == 0 {
		return nil
	}
	size := 1
	for _, item := range items {
		size *= len(item)
	}
	result := make([][]int32, 0, size)
	data := make([]int32, len(items))
	var backtrack func(int)
	backtrack = func(index int) {
		if len(items) == index {
			d := make([]int32, len(items))
			copy(d, data)
			result = append(result, d)
			return
		}
		for i := 0; i < len(items[index]); i++ {
			data[index] = items[index][i]
			backtrack(index + 1)
		}
	}
	backtrack(0)
	return result
}

// Map 对切片中的每个元素执行指定的函数
func Map[T any, Y any](data []T, fn func(int, T) Y) []Y {
	result := make([]Y, len(data))
	for i, v := range data {
		result[i] = fn(i, v)
	}
	return result
}

// Each 遍历数据
func Each[T any](data []T, fn func(int, T)) {
	for i, v := range data {
		fn(i, v)
	}
}

// Reduce 用回调函数迭代地将切片简化为单一的值
func Reduce[T any, Y any](data []T, fn func(int, T, Y) Y, init Y) Y {
	var result Y
	result = init
	for i, v := range data {
		result = fn(i, v, result)
	}
	return result
}

func CopyS[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)

	return dst
}

// SafeGet 安全的获取数组指定下标的值
func SafeGet[T any](arr []T, i int, missing T) T {
	if i >= len(arr) {
		return missing
	}

	return arr[i]
}

func Chunk[T any](arr []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	if len(arr) == 0 {
		return [][]T{}
	}

	var chunks [][]T
	for size < len(arr) {
		arr, chunks = arr[size:], append(chunks, arr[0:size])
	}
	return append(chunks, arr)
}

// Intersection 返回两个切片的交集
func Intersect[T comparable](a, b []T) []T {
	// 使用 map 记录第一个切片的元素
	set := make(map[T]struct{})
	for _, item := range a {
		set[item] = struct{}{}
	}

	// 检查第二个切片的元素是否在 map 中
	var result []T
	seen := make(map[T]struct{}) // 避免重复添加
	for _, item := range b {
		if _, exists := set[item]; exists {
			if _, added := seen[item]; !added {
				result = append(result, item)
				seen[item] = struct{}{}
			}
		}
	}
	return result
}

// Difference 返回第一个切片相对于第二个切片的差集
func Difference[T comparable](a, b []T) []T {
	// 使用 map 记录第二个切片的元素
	setB := make(map[T]struct{})
	for _, item := range b {
		setB[item] = struct{}{}
	}

	// 检查第一个切片的元素是否不在第二个切片中
	var result []T
	seen := make(map[T]struct{}) // 避免重复添加
	for _, item := range a {
		if _, exists := setB[item]; !exists {
			if _, added := seen[item]; !added {
				result = append(result, item)
				seen[item] = struct{}{}
			}
		}
	}
	return result
}

// Union 返回两个切片的并集
func Union[T comparable](a, b []T) []T {
	// 使用 map 记录所有唯一元素
	set := make(map[T]struct{})

	// 添加第一个切片的元素
	for _, item := range a {
		set[item] = struct{}{}
	}

	// 添加第二个切片的元素
	for _, item := range b {
		set[item] = struct{}{}
	}

	// 将 map 的键转换为切片
	result := make([]T, 0, len(set))
	for item := range set {
		result = append(result, item)
	}

	return result
}
