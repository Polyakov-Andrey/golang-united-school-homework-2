package square

import (
	"math"
	"testing"
)

const MESSAGE = "got %f, wanted %f"

func TestCalcSquareSidesTriangle(t *testing.T) {

	sideLen := 4.0
	got := CalcSquare(sideLen, SidesTriangle)
	want := sideLen * math.Sqrt(sideLen*sideLen-math.Pow(sideLen/2, 2)) / 2

	if got != want {
		t.Errorf(MESSAGE, got, want)
	}
}

func TestCalcSquareSidesSquare(t *testing.T) {

	sideLen := 4.0
	got := CalcSquare(sideLen, SidesSquare)
	want := math.Pow(sideLen, 2)

	if got != want {
		t.Errorf(MESSAGE, got, want)
	}
}

func TestCalcSquareSidesCircle(t *testing.T) {

	radius := 4.0
	got := CalcSquare(radius, SidesCircle)
	want := math.Pi * math.Pow(radius, 2)

	if got != want {
		t.Errorf(MESSAGE, got, want)
	}
}

func TestCalcSquareAny(t *testing.T) {

	sideLen := 4.0
	got := CalcSquare(sideLen, 5)
	want := 0.0

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
