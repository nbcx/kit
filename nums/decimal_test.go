package nums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimal(t *testing.T) {

	// 2 x (0.3 + 0.4) - 2 x 3
	dec := NewDecimal(2).Mul(0.3, 0.4).Sub(NewDecimal(2).Mul(3))
	assert.Equal(t, int64(-4), dec.Int64())
	assert.Equal(t, float32(-4.6), dec.Float32())

	// 1+1.7+1-0.3
	dec = NewDecimal(1)
	dec.Add(1.7)
	dec.Add(1)
	dec.Sub(0.3)
	assert.Equal(t, float64(3.4), dec.Float())
	assert.Equal(t, "3.4", dec.String())
}
