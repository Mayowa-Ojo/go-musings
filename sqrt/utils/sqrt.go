package utils

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %f", float64(e))
}

// Sqrt - Finds the square-root of a given int
func Sqrt(x float64) (float64, int, error) {
	var err ErrNegativeSqrt
	if x < 0 {
		return 0, 0, err
	}

	z := float64(1)
	itr := 0

	for itr < 10 {
		temp := z

		if temp -= adjustPrecision(temp, x); roundFloat(z, 11) == roundFloat(temp, 11) {
			// output, _ := fmt.Printf("square-root: %f - iterations: %v", z, itr)
			// return output
			return z, itr, nil
		}

		z -= adjustPrecision(z, x)
		itr++
	}
	// output, _ := fmt.Printf("square-root: %f - iterations: %v", z, itr)
	// return output
	return z, itr, nil

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
