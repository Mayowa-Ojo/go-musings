package main

import (
	"fmt"

	"github.com/Mayowa-Ojo/go-musings/sqrt/utils"
)

func main() {
	sqrt, iterations := utils.Sqrt(4)

	fmt.Printf("square-root: %f | iterations: %d", sqrt, iterations)
}
