package utils

import (
	"math"
)

// Sqrt - Finds the square-root of a given int
func Sqrt(x float64) (float64, int) {
	z := float64(1)
	itr := 0

	for itr < 10 {
		temp := z

		if temp -= adjustPrecision(temp, x); roundFloat(z, 11) == roundFloat(temp, 11) {
			// output, _ := fmt.Printf("square-root: %f - iterations: %v", z, itr)
			// return output
			return z, itr
		}

		z -= adjustPrecision(z, x)
		itr++
	}
	// output, _ := fmt.Printf("square-root: %f - iterations: %v", z, itr)
	// return output
	return z, itr

}

func adjustPrecision(v, x float64) float64 {
	return (v*v - x) / (2 * v)
}

func roundFloat(f float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * f
	round = math.Ceil(digit)
	return round / pow
}
