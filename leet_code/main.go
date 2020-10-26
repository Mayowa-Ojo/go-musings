package main

import (
	"fmt"

	"github.com/Mayowa-Ojo/go-musings/leet_code/lib"
)

func main() {
	testStr := "(()(()))"
	result := lib.MaxDepth(testStr)

	fmt.Printf("result: %d", result)
}
