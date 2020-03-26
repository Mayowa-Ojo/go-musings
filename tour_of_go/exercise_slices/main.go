package main

import "golang.org/x/tour/pic"

func main() {
	pic.Show(Pic)
}

/**
generator functions
	one: x*y
	two: (x+y)/2
	three: x^y
	four: 4(x^2 + y^2)
*/

// Pic - main function
func Pic(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	// Do something.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			val := generate(i, j, "four")
			a[i][j] = val
		}
	}
	return a
}

func generate(x, y int, generator string) uint8 {
	if generator == "one" {

		res := uint8(x * y)
		return res
	}

	if generator == "two" {
		res := uint8((x + y) / 2)
		return res
	}

	if generator == "three" {
		res := uint8(x ^ y)
		return res
	}

	if generator == "four" {
		res := uint8(4 * (x ^ 2 + y ^ 2))
		return res
	}

	return uint8(x + y)
}
