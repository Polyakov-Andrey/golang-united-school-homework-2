package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
import (
	"math"
)

type sides int

const (
	SidesTriangle sides   = 3
	SidesSquare   sides   = 4
	SidesCircle   sides   = 0
	Pi            float64 = math.Pi
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	var s float64
	switch sidesNum {
	case SidesTriangle:
		var p float64 = (3 * sideLen) / 2
		s = math.Sqrt(p * (p - sideLen) * (p - sideLen) * (p - sideLen))
	case SidesSquare:
		s = sideLen * sideLen
	case SidesCircle:
		s = Pi * sideLen * sideLen
	default:
		s = 0
	}
	return s
}
