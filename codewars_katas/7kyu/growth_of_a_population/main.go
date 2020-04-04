package main

import "fmt"

func main() {
	test1 := nbYear(1500, 5, 100, 5000)           // -> 15
	test2 := nbYear(1500000, 2.5, 10000, 2000000) // -> 10
	test3 := nbYear(1500000, 0.25, 1000, 2000000) // -> 94

	fmt.Printf("- %f\n- %f\n- %f", test1, test2, test3)
}

func nbYear(p0, percent, aug, p float64) (n float64) {
	for n = 1; n < p; n++ {
		p0 = p0 + (p0 * (percent / 100)) + aug

		if p0 >= p {
			break
		}
	}
	return n
}
