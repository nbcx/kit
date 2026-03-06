package nums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToString(t *testing.T) {
	result := IntToString(123)
	if result != "123" {
		t.Errorf("IntToString(123) failed. Got %s", result)
	}
}

func TestToStr(t *testing.T) {
	assert.Equal(t, "123", ToStr(int16(123)))
	assert.Equal(t, "-123", ToStr(-123))
	assert.Equal(t, "123.5", ToStr(123.5))
	assert.Equal(t, "-123.56", ToStr(-123.56))
}
