package main

import (
	"fmt"

	"github.com/Mayowa-Ojo/go-musings/sqrt/utils"
)

func main() {
	sqrt, iterations, err := utils.Sqrt(-4)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("square-root: %f | iterations: %d", sqrt, iterations)
}
