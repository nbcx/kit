package s

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveValues(t *testing.T) {
	type args struct {
		array []int32
		elem  []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"", args{[]int32{1, 2, 2, 3, 3, 5, 7}, []int32{2, 3, 3}}, []int32{1, 5, 7}},
		{"", args{[]int32{1, 2, 3, 4, 5, 6}, []int32{1, 2, 3}}, []int32{4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveValues(tt.args.array, tt.args.elem...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{[]int{1, 2, 3}, 1}, true},
		{"test2", args{[]int{1, 2, 3}, 3}, true},
		{"test2", args{[]int{1, 2, 3}, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	type args struct {
		data   []int
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{[]int{1, 2, 3}, []int{1, 2}}, true},
		{"test2", args{[]int{1, 2, 3}, []int{1, 3}}, true},
		{"test3", args{[]int{1, 2, 3}, []int{1, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.args.data, tt.args.values...); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindValueIndex(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3}, 1}, 0},
		{"test2", args{[]int{1, 2, 3}, 2}, 1},
		{"test3", args{[]int{1, 2, 3}, 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindValueIndex(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("FindValueIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPick(t *testing.T) {
	type args struct {
		data    []int
		indexes []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6}, []int{0, 1, 2}}, []int{1, 2, 3}},
		{"test2", args{[]int{1, 2, 3, 4, 5, 6}, []int{0, 3, 100}}, []int{1, 4}},
		{"test3", args{[]int{1, 2, 3, 4, 5, 6}, []int{4, 1}}, []int{5, 2}},
		{"test4", args{[]int{1, 2, 3, 4, 5, 6}, []int{50, 100, 200}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pick(tt.args.data, tt.args.indexes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		data  []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6}, 0}, []int{2, 3, 4, 5, 6}},
		{"test2", args{[]int{1, 2, 3, 4, 5, 6}, 1}, []int{1, 3, 4, 5, 6}},
		{"test2", args{[]int{1, 2, 3, 4, 5, 6}, 100}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.data, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6}}, []int{6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Reverse(tt.args.data); !reflect.DeepEqual(tt.args.data, tt.want) {
				t.Errorf("Reverse() = %v, want %v", tt.args.data, tt.want)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	Shuffle(data)
	assert.NotEqual(t, data, []int{1, 2, 3, 4, 5, 6})
}

func TestSum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6}}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.data); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEach(t *testing.T) {
	data := []int{1, 2, 3, 4}
	var sum int
	Each(data, func(i, num int) {
		sum += num
	})
	assert.Equal(t, sum, 10)
}

func TestMap(t *testing.T) {
	type args struct {
		data []int
		fn   func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"double", args{[]int{1, 2, 3, 4}, func(i, item int) int { return item * 2 }}, []int{2, 4, 6, 8}},
		{"square", args{[]int{1, 2, 3, 4}, func(i, item int) int { return item * item }}, []int{1, 4, 9, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Map(tt.args.data, tt.args.fn), "Map(%v, %v)", tt.args.data, tt.args.fn)
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		data []int
		fn   func(int, int, int) int
		init int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sum", args{[]int{1, 2, 3, 4}, func(i int, j int, x int) int { return j + x }, 0}, 10},
		{"product", args{[]int{1, 2, 3, 4}, func(i int, j int, x int) int { return j * x }, 1}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Reduce(tt.args.data, tt.args.fn, tt.args.init), "Reduce(%v, %v, %v)", tt.args.data, tt.args.fn, tt.args.init)
		})
	}
}

func TestRemoveFirstValue(t *testing.T) {
	type args struct {
		array []int32
		elem  int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"", args{[]int32{1, 2, 2, 3, 3, 5, 7}, 2}, []int32{1, 2, 3, 3, 5, 7}},
		{"", args{[]int32{1, 2, 3, 4, 5, 6}, 6}, []int32{1, 2, 3, 4, 5}},
		{"", args{[]int32{1}, 2}, []int32{1}},
		{"", args{[]int32{1}, 1}, []int32{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFirstValue(tt.args.array, tt.args.elem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirstValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyS(t *testing.T) {
	srcInt := []int{1, 2, 3}
	dstInt := CopyS(srcInt)
	assert.Equal(t, srcInt, dstInt)

	srcFloat := []float32{0.1, 0.2, 0.3}
	dstFloat := CopyS(srcFloat)
	assert.Equal(t, srcFloat, dstFloat)
}

func TestSafeGet(t *testing.T) {
	type args[T any] struct {
		arr     []T
		i       int
		missing T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "t1",
			args: args[int]{
				arr:     []int{1, 2, 3, 4},
				i:       1,
				missing: 0,
			},
			want: 2,
		},
		{
			name: "t2",
			args: args[int]{
				arr:     []int{1, 2, 3, 4},
				i:       100,
				missing: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SafeGet(tt.args.arr, tt.args.i, tt.args.missing), "SafeGet(%v, %v, %v)", tt.args.arr, tt.args.i, tt.args.missing)
		})
	}
}

func TestChunk(t *testing.T) {
	type args struct {
		arr  []int
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Chunk with size less than array length",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 3},
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
		},
		{
			name: "Chunk with size greater than array length",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 15},
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		},
		{
			name: "Chunk with size equal to array length",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 10},
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		},
		{
			name: "Chunk with size zero",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 0},
			want: nil,
		},
		{
			name: "Chunk with empty array",
			args: args{arr: []int{}, size: 3},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.arr, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"test1", []int{1, 4, 7}, []int{1, 4}, []int{1, 4}},
		{"test2", []int{1, 4, 7}, []int{2, 5}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersect(tt.a, tt.b)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"test1", []int{1, 4, 7}, []int{1, 4}, []int{7}},
		{"test2", []int{1, 4, 7}, []int{2, 5}, []int{1, 4, 7}},
		{"test3", []int{1, 4, 7}, []int{1, 4, 7}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.a, tt.b)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"test1", []int{1, 4, 7}, []int{1, 4}, []int{1, 4, 7}},
		{"test2", []int{1, 4, 7}, []int{2, 5}, []int{1, 2, 4, 5, 7}},
		{"test3", []int{1, 4, 7}, []int{1, 4, 7}, []int{1, 4, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Union(tt.a, tt.b)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsEqual(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		{"test1", []int{1, 4}, []int{1, 4, 5}, false},
		{"test1", []int{1, 4}, []int{4, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("ContainsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
