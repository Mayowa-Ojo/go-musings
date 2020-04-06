package main

import (
	"fmt"
	"strings"
)

func main() {
	res := Accum("abcd")

	fmt.Println(res)
}

// Accum - repeat string element in increasing order of index
func Accum(s string) string {
	s = strings.ToLower(s)
	buildStr := ""

	for i := 0; i < len(s); i++ {
		rptStr := strings.ToUpper(string(s[i])) + strings.Repeat(string(s[i]), i) + "-"

		buildStr += rptStr
	}

	r := []rune(buildStr)
	subStr := string(r[0 : len(buildStr)-1])
	return subStr
}
