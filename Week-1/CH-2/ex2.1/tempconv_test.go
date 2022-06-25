package tempconv

import (
	"math"
	"testing"
)

func TestTempConv(t *testing.T) {
	tests := []struct {
		f Fahrenheit
	}{
		{68, 20, 293.15},
		{32, 0, 273.15},
		{-40, -40, 233.15},
	}
	eps := 0.0000001 // acceptable floating point error
	for _, test := range tests {
		
		if math.Abs(float64(FToC(test.f)-test.c)) > eps {
			t.Errorf("FToC(%s): got %s, want %s", test.f, FToC(test.f), test.c)
		}
		
	}
}
