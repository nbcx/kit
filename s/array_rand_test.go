package s

import (
	"math"
	"testing"

	"github.com/nbcx/kit/nums"
	"github.com/stretchr/testify/assert"
)

func TestRandChoice(t *testing.T) {
	type args struct {
		array []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6}, 3}, 3},
		{"test1", args{[]int{1, 2, 3, 4, 5, 6}, 100}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			randData := RandChoice(tt.args.array, tt.args.n)
			assert.Equalf(t, tt.want, len(randData), "RandChoice(%v, %v)", tt.args.array, tt.args.n)
		})
	}
}

func TestWeightedChoiceFloat32(t *testing.T) {
	type args struct {
		weights []float32
		loop    int
	}
	tests := []struct {
		name string
		args args
		want float64 // 误差
	}{
		{"test1", args{[]float32{1.0, 2.0, 3.0, 4.0}, 10000 * 100}, 0.01},
		{"test2", args{[]float32{150.0, 200.0, 240.0, 300.0, 400.0}, 10000 * 100}, 0.01},
		{"test3", args{[]float32{0.0, 0.0, 0.0, 300.0, 400.0}, 10000 * 100}, 0.01},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countMap := make([]int, len(tt.args.weights))
			for i := 0; i < tt.args.loop; i++ {
				index, _ := WeightedChoiceFloat32(tt.args.weights)
				countMap[index]++
			}
			theories := make([]float64, len(tt.args.weights))
			for i, weight := range tt.args.weights {
				theories[i] = float64(weight / Sum(tt.args.weights))
			}

			for i := 0; i < len(countMap); i++ {
				practice := float64(countMap[i]) / float64(tt.args.loop)
				if nums.IsEqualFloat64(theories[i], 0.0) {
					assert.Equal(t, true, nums.IsEqualFloat64(theories[i], practice))
				} else {
					assert.GreaterOrEqualf(t, tt.want, math.Abs(theories[i]-practice)/theories[i], "WeightedChoiceFloat32(%v)", tt.args.weights)
				}
			}
		})
	}
}

func TestPerm(t *testing.T) {
	tests := []struct {
		name  string
		total int
		n     int
	}{
		{"test1", 5000, 100},
		{"test2", 10, 10},
		{"test3", 10, 1},
		{"test4", 10000, 2},
		{"test5", 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Perm[int](tt.total, tt.n)
			if len(got) != tt.n {
				t.Errorf("Perm() = %v, want %v", got, tt.n)
			}

			var result = make(map[int]struct{}, tt.n)
			for _, v := range got {
				if _, ok := result[v]; ok {
					t.Errorf("Perm() = %v", got)
				}
			}
		})
	}
}

func TestWeightedChoice(t *testing.T) {
	got, _ := Weighted([]int64{23, 0, 0, 0})
	if got != 0 {
		t.Errorf("WeightedChoice() = %v, want %v", got, 0)
	}

	_, err := Weighted([]int64{0, 0, 0, 0})
	if err == nil {
		t.Errorf("WeightedChoice() = %v, want %v", err, "no nil")
	}

	_, err = Weighted([]int64{})
	if err == nil {
		t.Errorf("WeightedChoice() = %v, want %v", err, "no nil")
	}

}

func TestCombinations(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		m    int32
		want int
	}{
		{"test1", []int32{1, 2, 3, 4}, 3, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combinations(tt.arr, tt.m); len(got) != tt.want {
				t.Errorf("Combinations() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
