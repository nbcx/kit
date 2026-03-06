package nums

import (
	"testing"
)

func TestRandFloat(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := RandFloat(1.3, 2.5, 1)
		if num < 1.3 || num > 2.5 {
			t.Errorf("RandFloat error: %f", num)
		}
	}
}

func TestRand(t *testing.T) {

	num := Rand(0.3, 1)
	if num < 0.3 || num > 1 {
		t.Errorf("rand error: %f", num)
	}

	num2 := Rand(0.3, 0.9)
	if num2 < 0.3 || num2 > 0.9 {
		t.Errorf("rand error: %f", num2)
	}

	num3 := Rand(0.3, 0.3)
	if num3 < 0.3 || num3 > 0.3 {
		t.Errorf("rand error: %f", num3)
	}

	num4 := Rand(3, 6)
	if num4 < 3 || num4 > 6 {
		t.Errorf("rand error: %d", num4)
	}

	num5 := Rand(1, 1)
	if num5 < 1 || num5 > 1 {
		t.Errorf("rand error: %d", num5)
	}
}
