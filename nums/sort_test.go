package nums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	nums := []int64{10, 3, 5, 8, 6}
	Sort(nums)
	assert.Equal(t, int64(3), nums[0])

	floats := []float32{10.1, 3.1, 3.2, 8.8, 10.1}
	Sort(floats)
	assert.Equal(t, float32(10.1), floats[4])

	type MyInt int
	my := []MyInt{6, 4, 5}
	Sort(my)
	assert.Equal(t, MyInt(6), my[2])
}
