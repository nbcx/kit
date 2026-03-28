package f

import (
	"fmt"
	"testing"
	"time"
)

func TestSpeed(t *testing.T) {
	tests := []struct {
		name string
		fs   []func()
		want time.Duration
	}{
		{"test", []func(){func() { time.Sleep(time.Second) }}, time.Second},
		{"test", []func(){func() { time.Sleep(time.Second) }, func() { time.Sleep(2 * time.Second) }}, 3 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Speed(tt.fs...); got < tt.want {
				t.Errorf("Speed() = %v, want %v", got, tt.want)
			} else {
				fmt.Println("got", got, tt.want, got < tt.want)
			}
		})
	}
}

func TestIsNil(t *testing.T) {

	type Inter interface{}

	var u Inter

	type Child struct{}

	var c *Child
	var cc Inter

	cc = c
	fmt.Printf("%+v,%+v\n", c, Inter(c) == nil)
	var b any = (*int)(nil)
	tests := []struct {
		name  string
		args  any
		want1 bool
		want2 bool
	}{
		{"test1", nil, true, true},
		{"test2", false, false, false},
		{"test3", u, true, true},
		{"test4", cc, false, true},
		{"test5", c, false, true},
		{"test6", b, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (tt.args == nil); got != tt.want1 {
				t.Errorf("nil = %v, want1 %v", got, tt.want1)
			}
			if got := IsNil(tt.args); got != tt.want2 {
				t.Errorf("IsNil() = %v, want2 %v", got, tt.want2)
			}
		})
	}
}
