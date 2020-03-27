package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	hold := 0
	prevVal := 0
	nextVal := 1
	return func() int {
		if count == 0 {
			count++
			return 0
		}

		if count == 1 {
			count++
			return 1
		}

		hold = nextVal
		nextVal = prevVal + nextVal
		prevVal = hold
		return nextVal
	}
}

// shorter version
func optFibonacci() func() int {
	prev, nextv := 0, 1

	return func() int {
		res := prev
		prev, nextv = nextv, prev+nextv
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
