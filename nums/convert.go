package nums

import (
	"fmt"
	"reflect"

	"github.com/nbcx/kit/types"
)

// ConvertToInt guess Num format and convert to Int
func ConvertToInt(temp interface{}) (int, error) {
	switch t := temp.(type) {
	case int:
		return t, nil
	case float64, float32:
		return int(reflect.ValueOf(t).Float()), nil
	case int64, int32:
		return int(reflect.ValueOf(t).Int()), nil
	default:
		return 0, fmt.Errorf("can't convert to int:%v", temp)
	}
}

var floatType = reflect.TypeOf(float64(0))

// ConvertToFloat64 guess Num format and convert to Float64
func ConvertToFloat64(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

// Float32ToInt Float32ToInt
func Float32ToInt(f float32) int {
	return int(f + 0.5)
}

// Float64ToInt Float64ToInt
func Float64ToInt(f float64) int {
	return int(f + 0.5)
}

// Float32ToInt32 Float32ToInt32
func Float32ToInt32(f float32) int32 {
	return int32(f + 0.5)
}

// Float64ToInt32 Float64ToInt32
func Float64ToInt32(f float64) int32 {
	return int32(f + 0.5)
}

// Float32ToInt64 Float32ToInt64
func Float32ToInt64(f float32) int64 {
	return int64(f + 0.5)
}

// Float64ToInt64 Float64ToInt64
func Float64ToInt64(f float64) int64 {
	return int64(f + 0.5)
}

// IntToString IntToString
func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

func ToStr[T types.Number](n T) string {
	return fmt.Sprintf("%v", n)
}
