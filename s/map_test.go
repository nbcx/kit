package s

import (
	"sync"
	"testing"

	"github.com/nbcx/kit/nums"
	"github.com/nbcx/kit/to"
	"github.com/stretchr/testify/assert"
)

func TestCountSyncMap(t *testing.T) {
	var data sync.Map
	data.Store(1, 1)
	data.Store(2, 2)
	assert.Equal(t, 2, CountSyncMap(&data))
}

func TestMapKeys(t *testing.T) {
	type args struct {
		data map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{map[string]int{"a": 1, "b": 2, "c": 3}}, 3},
		{"test2", args{map[string]int{"qq": 1, "ff": 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := len(MapKeys(tt.args.data))
			assert.Equalf(t, tt.want, count, "MapKeys(%v)", tt.args.data)
		})
	}
}

func TestHasKey(t *testing.T) {
	type args struct {
		data map[string]int
		key  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{map[string]int{"a": 1, "b": 2, "c": 3}, "a"}, true},
		{"test2", args{map[string]int{"qq": 1, "ff": 2, "a": 3}, "ff"}, true},
		{"test3", args{map[string]int{"qq": 1, "ff": 2, "a": 3}, "d"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MapHasKey(tt.args.data, tt.args.key), "HasKey(%v, %v)", tt.args.data, tt.args.key)
		})
	}
}

func TestMapToMap(t *testing.T) {
	r1 := MapToMap(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}, func(k string, v int) (string, int) {
		return k + "1", v + 1
	})

	assert.Equal(t, map[string]int{
		"a1": 2,
		"b1": 3,
		"c1": 4,
	}, r1)

	r2 := MapToMap(map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}, func(k string, v int) (int, string) {
		return to.Int(k), nums.IntToString(v)
	})
	assert.Equal(t, map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}, r2)
}

func TestMapEach(t *testing.T) {
	var sum int
	MapEach(map[string]int{
		"aa": 1,
		"bb": 2,
		"cc": 3,
	}, func(_ string, v int) {
		sum += v
	})
	assert.Equal(t, 6, sum)
}

func TestMapReduce(t *testing.T) {
	sum := MapReduce(map[string]int{
		"aa": 1,
		"bb": 2,
		"cc": 3,
	}, func(_ string, v int, sum int) int {
		return sum + v
	}, 0)
	assert.Equal(t, 6, sum)
}

func TestMapKeysSum(t *testing.T) {
	sum := MapKeysSum(map[int]string{
		1: "aa",
		2: "bb",
		3: "cc",
	})
	assert.Equal(t, 6, sum)
}

func TestMapValuesSum(t *testing.T) {
	sum := MapValuesSum(map[string]int{
		"aa": 1,
		"bb": 2,
		"cc": 3,
	})
	assert.Equal(t, 6, sum)
}

func TestMapMerge(t *testing.T) {
	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	map2 := map[string]int{
		"a": 11,
		"c": 3,
		"d": 4,
	}

	map3 := map[string]int{
		"e": 5,
		"f": 6,
		"c": 33,
	}

	r := MapMerge(map1, map2, map3, nil)
	assert.Equal(t, map[string]int{
		"a": 11,
		"b": 2,
		"c": 33,
		"d": 4,
		"e": 5,
		"f": 6,
	}, r)
}
