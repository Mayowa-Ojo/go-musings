package main

import (
	"fmt"
	. "github.com/mayowa-ojo/foreach/module"
)

func main() {
	names := CustomArray{"Joe", "Henry", "Joy", "Amy"}
	names.Foreach(print)

}

func print(element string, index int) {
	fmt.Println(element, index)
}