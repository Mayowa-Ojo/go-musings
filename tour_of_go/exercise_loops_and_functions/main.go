package main

import "fmt"

func main() {
	fmt.Println(Sqrt(4))
}

// Sqrt - evaluate sqrt of a floating point number
func Sqrt(x float64) (float64, int) {
	z := float64(1)
	itr := 0

	for itr < 10 {
		temp := z
		if temp -= (temp*temp - x) / (2 * temp); z == temp {
			return z, itr
		}
		z -= (z*z - x) / (2 * z)

		itr++
	}
	return z, itr
}
