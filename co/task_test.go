package co

import (
	"reflect"
	"testing"
	"time"

	"github.com/nbcx/kit/k"
)

func TestTask(t *testing.T) {
	total := 0
	test := func() int {
		time.Sleep(time.Second * 3)
		return 1
	}
	call := WithCallback(func(t int) {
		total += t
	})
	speed := k.Speed(func() {
		t := NewTask(call)
		for i := 0; i < 5; i++ {
			t.Exec(test)
		}
		t.Wait()
	})
	if !reflect.DeepEqual(total, 5) {
		t.Errorf("total(%v) !=%v", total, 5)
	}
	t.Log("speed", speed)
}
