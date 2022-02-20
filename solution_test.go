package square

import (
	"math"
	"testing"
)

func TestCalcSquare_SidesTriangle(t *testing.T) {

	sideLen := 4.0
	got := CalcSquare(sideLen, SidesTriangle)
	want := sideLen * math.Sqrt(sideLen*sideLen-math.Pow(sideLen/2, 2)) / 2

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestCalcSquare_SidesSquare(t *testing.T) {

	sideLen := 4.0
	got := CalcSquare(sideLen, SidesSquare)
	want := math.Pow(sideLen, 2)

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestCalcSquare_SidesCircle(t *testing.T) {

	sideLen := 4.0
	got := CalcSquare(sideLen, SidesCircle)
	want := math.Pi * math.Pow(sideLen, 2)

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestCalcSquare_Any(t *testing.T) {

	sideLen := 4.0
	got := CalcSquare(sideLen, 5)
	want := 0.0

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
