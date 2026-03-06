package nums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqualFloat32(t *testing.T) {
	type args struct {
		x float32
		y float32
	}
	tests := []struct {
		name string
		args args
		want bool // 是否相等，true相等，false不等
	}{
		{"test1", args{0.0, 0.0}, true},
		{"test2", args{1.0, 1.0}, true},
		{"test3", args{1e-5, 1e-5}, true},
		{"test4", args{1e-6, 1e-6}, true},
		{"test5", args{1e-7, 1e-7}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsEqualFloat32(tt.args.x, tt.args.y), "IsEqual(%v)", tt.args)
		})
	}
}

func TestIsEqualFloat64(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want bool // 是否相等，true相等，false不等
	}{
		{"test1", args{0.0, 0.0}, true},
		{"test2", args{1.0, 1.0}, true},
		{"test3", args{1e-14, 1e-14}, true},
		{"test4", args{1e-15, 1e-15}, true},
		{"test5", args{1e-16, 1e-16}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsEqualFloat64(tt.args.x, tt.args.y), "IsEqual(%v)", tt.args)
		})
	}
}

func TestIsBetween(t *testing.T) {
	type args struct {
		val     int
		between []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{1, []int{0, 2}}, true},
		{"test2", args{1, []int{1, 2}}, true},
		{"test3", args{1, []int{0, 1}}, true},
		{"test4", args{1, []int{0, 0}}, false},
		{"test5", args{1, []int{2, 8}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsBetween(tt.args.val, tt.args.between), "IsBetween(%v, %v)", tt.args.val, tt.args.between)
		})
	}
}
