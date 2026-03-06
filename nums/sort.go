package nums

import (
	"sort"

	"github.com/nbcx/kit/types"
)

// IntSlice .
type IntSlice[T types.Number] []T

func (x IntSlice[T]) Len() int           { return len(x) }
func (x IntSlice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x IntSlice[T]) Sort() { Sort(x) }

// Sort sorts data in ascending order as determined by the Less method.
func Sort[T types.Number](x []T) {
	sort.Sort(IntSlice[T](x))
}
