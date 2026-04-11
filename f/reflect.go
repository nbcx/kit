package f

import "reflect"

// Ptr returns a pointer of T.
// useful when you want to get a pointer of a value, in some config, for example.
// eg. Ptr[int] returns *int.
// eg. Ptr[*int] returns **int.
func Ptr[T any](v T) *T {
	return &v
}

// TypeOf returns the type of T.
// eg. TypeOf[int] returns reflect.TypeOf(int).
// eg. TypeOf[*int] returns reflect.TypeOf(*int).
func TypeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

// NewInstance create an instance of the given type T.
// the main purpose of this function is to create an instance of a type, can handle the type of T is a pointer or not.
// eg. NewInstance[int] returns 0.
// eg. NewInstance[*int] returns *0 (will be ptr of 0, not nil!).
func NewInstance[T any]() T {
	typ := TypeOf[T]()
	switch typ.Kind() {
	case reflect.Map:
		return reflect.MakeMap(typ).Interface().(T)
	case reflect.Slice, reflect.Array:
		return reflect.MakeSlice(typ, 0, 0).Interface().(T)
	case reflect.Ptr:
		typ = typ.Elem()
		origin := reflect.New(typ)
		inst := origin

		for typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			inst = inst.Elem()
			inst.Set(reflect.New(typ))
		}

		return origin.Interface().(T)
	default:
		var t T
		return t
	}
}
