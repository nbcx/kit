package f

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	fmt.Println(If(1 > 2, 1, 2))
}

func TestIf(t *testing.T) {
	type args struct {
		expr bool
		a    any
		b    any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{"int", args{1 > 2, 1, 2}, 2},
		{"string", args{"hello" != "world", "hello", "world"}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.expr, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}
