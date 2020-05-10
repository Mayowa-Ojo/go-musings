package main

import (
	"math"
	"strconv"
	"fmt"
	"sort"
)

func main() {
	res := primeFactors(7775460)
	fmt.Println(res)
}

func PrimeFactors(n int) string {
	var result string
	sortedKeys := make([]int, 0)

	if n == 0 || n == 1 {
		return "(" + strconv.Itoa(n) + ")"
	}

	m := make(map[int]int)

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if _, ok := m[i]; !ok {
				m[i] = 0
			}
			m[i] = m[i] + 1
			n = int(math.Floor(float64(n) / float64(i)))
		}
	}

	// sort map
	for i, _ := range m {
		sortedKeys = append(sortedKeys, i)
	}

	sort.Ints(sortedKeys)
	// format result
	for _, v := range sortedKeys {
		if m[v] > 1 {
			result += ("(" + strconv.Itoa(v) + "**" + strconv.Itoa(m[v]) + ")")
      			continue
		}
		result += ("(" + strconv.Itoa(v) + ")")
	}

	return result
}

