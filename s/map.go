package s

import (
	"sync"

	"github.com/nbcx/kit/types"
)

func CountSyncMap(m *sync.Map) int {
	size := 0
	m.Range(func(_, _ interface{}) bool {
		size++
		return true
	})
	return size
}

// MapKeys 返回map的所有key
func MapKeys[K comparable, V any](data map[K]V) []K {
	keys := make([]K, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}

// MapHasKey map是否包含指定的key
func MapHasKey[K comparable, V any](data map[K]V, key K) bool {
	_, ok := data[key]
	return ok
}

// MapToMap map转map
func MapToMap[K1 comparable, V1 any, K2 comparable, V2 any](data map[K1]V1, fn func(K1, V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(data))
	for k, v := range data {
		k2, v2 := fn(k, v)
		result[k2] = v2
	}
	return result
}

// MapEach 遍历map
func MapEach[K comparable, V any](data map[K]V, fn func(K, V)) {
	for k, v := range data {
		fn(k, v)
	}
}

// MapReduce map reduce
func MapReduce[K comparable, V any, R any](data map[K]V, fn func(K, V, R) R, init R) R {
	for k, v := range data {
		init = fn(k, v, init)
	}
	return init
}

// MapValuesSum 求所有value的和
func MapValuesSum[K comparable, V types.Number](data map[K]V) V {
	var sum V
	for _, v := range data {
		sum += v
	}
	return sum
}

// MapValuesSum 求所有key的和
func MapKeysSum[K types.Number, V any](data map[K]V) K {
	var sum K
	for k := range data {
		sum += k
	}
	return sum
}

// MapMerge 合并多个 map
func MapMerge[K comparable, V any](map1 map[K]V, mapN ...map[K]V) map[K]V {
	for _, item := range mapN {
		if item == nil {
			continue
		}

		for k, v := range item {
			map1[k] = v
		}
	}
	return map1
}
