package main

import (
	"fmt"
	"regexp"
)

func main() {

	match := validatePin("1234")

	fmt.Println(match)

}

func validatePin(pin string) bool {
	rgx, _ := regexp.MatchString(`^\d+$`, pin)
	isValid := (len(pin) == 4 || len(pin) == 6) && rgx

	return isValid
}
