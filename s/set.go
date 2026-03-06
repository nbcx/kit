package s

// Set 集合，非goroutine安全
type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	d := &Set[T]{
		set: make(map[T]struct{}, len(items)),
	}
	for _, item := range items {
		d.set[item] = struct{}{}
	}
	return d
}

func (d *Set[T]) Add(items ...T) *Set[T] {
	if d.set == nil {
		d.set = make(map[T]struct{}, len(items))
	}
	for _, item := range items {
		d.set[item] = struct{}{}
	}
	return d
}

func (d *Set[T]) Remove(items ...T) *Set[T] {
	for _, item := range items {
		delete(d.set, item)
	}
	return d
}

func (d *Set[T]) Contains(items ...T) bool {
	var ok bool
	for _, item := range items {
		if _, ok = d.set[item]; !ok {
			return false
		}
	}
	return true
}

func (d *Set[T]) Size() int {
	return len(d.set)
}

// Intersect 交集
func (d *Set[T]) Intersect(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	// 遍历较小的那个
	toRange, another := d.set, other
	if d.Size() > other.Size() {
		toRange, another = other.set, d
	}
	for k := range toRange {
		if another.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// Union 并集
func (d *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k, v := range d.set {
		result.set[k] = v
	}
	for k, v := range other.set {
		result.set[k] = v
	}
	return result
}

// Difference 差集
func (d *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k := range d.set {
		if !other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

func (d *Set[T]) ToArray() []T {
	result := make([]T, 0, d.Size())
	for k := range d.set {
		result = append(result, k)
	}
	return result
}
