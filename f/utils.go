package ext

import (
	"reflect"
	"time"
)

// ReTry 重试逻辑，最多重试n次
func ReTry(n int, f func() error) error {
	var err error
	for i := 0; i < n; i++ {
		err = f()
		if err == nil {
			return nil
		}
	}
	return err
}

// Speed 计算函数执行时间
func Speed(fs ...func()) time.Duration {
	start := time.Now()
	for _, v := range fs {
		v()
	}
	return time.Since(start)
}

// Copy 拷贝
func Copy[T any](v []T) []T {
	dst := make([]T, len(v))
	copy(dst, v)
	return dst
}

// IsNil 判读是否为空，主要是解决any类型的判断
func IsNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Interface, reflect.Func:
		return rv.IsNil()
	}
	return false
}
