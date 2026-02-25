package to

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Value struct {
	v any
}

func V(v any) *Value {
	return &Value{v}
}

func (s *Value) IsNil() bool {
	return s == nil || s.v == nil
}

// Int returns input as an int
func (v *Value) Int() int {
	i, _ := v.IntE()
	return i
}

func (v *Value) IntE() (i int, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(int); ok {
		return vv, nil
	}
	return strconv.Atoi(fmt.Sprint(v.v))
}

func (v *Value) Int8() int8 {
	i, _ := v.Int8E()
	return i
}

// Int8E return input as an int8
func (v *Value) Int8E() (i8 int8, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(int8); ok {
		return vv, nil
	}
	i64, err := strconv.ParseInt(fmt.Sprint(v.v), 10, 8)
	return int8(i64), err
}

// GetUint8 return input as an uint8
func (v *Value) Uint8() (i uint8) {
	i, _ = v.Uint8E()
	return
}

// GetUint8 return input as an uint8
func (v *Value) Uint8E() (u8 uint8, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(uint8); ok {
		return vv, nil
	}
	u64, err := strconv.ParseUint(fmt.Sprint(v.v), 10, 8)
	return uint8(u64), err
}

func (v *Value) Int16() (i int16) {
	i, _ = v.Int16E()
	return
}

// GetInt16 returns input as an int16
func (v *Value) Int16E() (i16 int16, e error) {
	if v.IsNil() {
		return
	}

	if vv, ok := v.v.(int16); ok {
		return vv, nil
	}
	i64, err := strconv.ParseInt(fmt.Sprint(v.v), 10, 16)
	return int16(i64), err
}

func (v *Value) Uint16() (i uint16) {
	i, _ = v.Uint16E()
	return
}

// Uint16E returns input as an uint16 and return error if has
func (v *Value) Uint16E() (u16 uint16, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(uint16); ok {
		return vv, nil
	}
	u64, err := strconv.ParseUint(fmt.Sprint(v.v), 10, 16)
	return uint16(u64), err
}

func (v *Value) Int32() (i int32) {
	i, _ = v.Int32E()
	return
}

// Int32E returns input as an int32
func (v *Value) Int32E() (i32 int32, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(int32); ok {
		return vv, nil
	}
	i64, err := strconv.ParseInt(fmt.Sprint(v.v), 10, 32)
	return int32(i64), err
}

func (v *Value) Uint32() (i uint32) {
	i, _ = v.Uint32E()
	return
}

// Uint32E returns input as an uint32
func (v *Value) Uint32E() (u32 uint32, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(uint32); ok {
		return vv, nil
	}
	u64, err := strconv.ParseUint(fmt.Sprint(v.v), 10, 32)
	return uint32(u64), err
}

func (v *Value) Int64() (i int64) {
	i, _ = v.Int64E()
	return
}

// Int64E returns input value as int64
func (v *Value) Int64E() (i64 int64, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(int64); ok {
		return vv, nil
	}
	return strconv.ParseInt(fmt.Sprint(v.v), 10, 64)
}

func (v *Value) Uint64() (i uint64) {
	i, _ = v.Uint64E()
	return
}

// Uint64E returns input value as uint64
func (v *Value) Uint64E() (u64 uint64, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(uint64); ok {
		return vv, nil
	}
	return strconv.ParseUint(fmt.Sprint(v.v), 10, 64)
}

func (v *Value) Bool() (i bool) {
	i, _ = v.BoolE()
	return
}

// GetBool returns input value as bool
func (v *Value) BoolE() (b bool, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(bool); ok {
		return vv, nil
	}
	return strconv.ParseBool(fmt.Sprint(v.v))
}

func (v *Value) Float() (i float64) {
	i, _ = v.FloatE()
	return
}

// FloatE returns input value as float64
func (v *Value) FloatE() (f float64, e error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(float64); ok {
		return vv, nil
	}
	return strconv.ParseFloat(fmt.Sprint(v.v), 64)
}

func (v *Value) Float32() (i float32) {
	i, _ = v.Float32E()
	return
}

// FloatE returns input value as float64
func (v *Value) Float32E() (f float32, err error) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(float32); ok {
		return vv, nil
	}
	vv, err := strconv.ParseFloat(fmt.Sprint(v.v), 32)
	return float32(vv), err
}

// Any returns input value as any
func (v *Value) Any() any {
	return v.v
}

// String returns input value as float64
func (v *Value) String() (s string) {
	if v.IsNil() {
		return
	}
	if vv, ok := v.v.(string); ok {
		return vv
	}
	return fmt.Sprint(v.v)
}

func (v *Value) Split(sep string) Values[string] {
	sp := strings.Split(v.String(), sep)
	return sp
}

func (v *Value) Int8s() (i Values[int8]) {
	i, _ = v.Int8sE()
	return
}

func (s *Value) array() *reflect.Value {
	switch reflect.TypeOf(s.v).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(s.v)
		return &s
		// var values = make([]int8, s.Len())

		// for i := 0; i < s.Len(); i++ {
		// 	values[i] = ValueF(s.Index(i)).Int8()
		// }
		// return Values[int8](values), nil
	}
	return nil
}

func (s *Value) Int8sE() (Values[int8], error) {
	if v, ok := s.v.([]int8); ok {
		return Values[int8](v), nil
	}
	v := s.array()
	if v == nil {
		return nil, fmt.Errorf("cant get []int")
	}
	var values = make([]int8, v.Len())
	for i := 0; i < v.Len(); i++ {
		values[i] = V(v.Index(i)).Int8()
	}
	return Values[int8](values), nil
}

func (v *Value) Ints() (i []int) {
	i, _ = v.IntsE()
	return
}

func (v *Value) VInts() (i Values[int]) {
	i, _ = v.IntsE()
	return
}

func (s *Value) IntsE() ([]int, error) {
	if v, ok := s.v.([]int); ok {
		return v, nil
	}
	v := s.array()
	if v == nil {
		return nil, fmt.Errorf("cant get []int")
	}
	var values = make([]int, v.Len())
	for i := 0; i < v.Len(); i++ {
		values[i] = V(v.Index(i)).Int()
	}
	return values, nil
}

func (s *Value) Values() (values []*Value) {
	anys, ok := s.v.([]any)
	if ok {
		for _, v := range anys {
			values = append(values, V(v))
		}
	}
	return
}

func (s *Value) ValuesE() (values []*Value, err error) {
	anys, ok := s.v.([]any)
	if ok {
		for _, v := range anys {
			values = append(values, V(v))
		}
		return
	}
	return nil, fmt.Errorf("cant get []*Value")
}

func (s *Value) ValueM() map[string]*Value {
	m := map[string]*Value{}
	anyM, ok := s.v.(map[string]any)
	if ok {
		for k, v := range anyM {
			m[k] = V(v)
		}
	}
	return m
}

func (s *Value) ValueME() (map[string]*Value, error) {
	m := map[string]*Value{}
	anyM, ok := s.v.(map[string]any)
	if ok {
		for k, v := range anyM {
			m[k] = V(v)
		}
		return m, nil
	}
	return m, fmt.Errorf("cant get map[string]*Value")
}

type Values[T any] []T

func Vs[T any](v []T) Values[T] {
	return Values[T](v)
}

func (v Values[T]) Int32() (i []int32) {
	for _, vv := range v {
		i = append(i, V(vv).Int32())
	}
	return
}

func (v Values[T]) Int() (i []int) {
	for _, vv := range v {
		i = append(i, V(vv).Int())
	}
	return
}

func (v Values[T]) String() (i []string) {
	for _, vv := range v {
		i = append(i, V(vv).String())
	}
	return
}

func (v Values[T]) First() (val *Value) {
	for _, vv := range v {
		return V(vv)
	}
	return
}

type ValueM[K comparable, T any] map[K]T

func Vm[K comparable, T any](v map[K]T) ValueM[K, T] {
	return ValueM[K, T](v)
}

func (v ValueM[K, T]) String() (i map[K]string) {
	i = make(map[K]string)
	for k, vv := range v {
		i[k] = V(vv).String()
	}
	return
}

func (v ValueM[K, T]) Int() (i map[K]int) {
	i = make(map[K]int)
	for k, vv := range v {
		i[k] = V(vv).Int()
	}
	return
}

func (v ValueM[K, T]) IntE() (i map[K]int) {
	i = make(map[K]int)
	for k, vv := range v {
		i[k] = V(vv).Int()
	}
	return
}

func (v ValueM[K, T]) KInt() (i map[int]T) {
	i = make(map[int]T)
	for k, vv := range v {
		i[V(k).Int()] = vv
	}
	return
}

func (v ValueM[K, T]) KVInt() (i map[int]int) {
	i = make(map[int]int)
	for k, vv := range v {
		i[V(k).Int()] = V(vv).Int()
	}
	return
}
