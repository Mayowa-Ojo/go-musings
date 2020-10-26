package lib

import (
	"sort"
	"strings"
)

// Problem | Maximun nesting depth of the parenthesis
// A string is a valid parenthesis string (denoted VPS) if it meets one of the following:
// 	- It is an empty string "", or a single character not equal to "(" or ")"
// 	- It can be written as AB (A concatenated with B), where A and B are VPS's, or
// 	- It can be written as (A), weher A is a VPS
// We can similarly define the nesting depth depth(S) of any VPS as follows:
// 	- depth("") = 0
// 	- depth(A + B) = max(depth(A), depth(B)), where A and B are VPS's
//		- depth("(" + A + ")") = 1 + depth(A), where A is a VPS
// For example, "", "()()", and "()(()())" are VPS's (with nesting depths 0, 1 and 2). and ")(" and "(()" are not VPS's.
//
// Given a VPS represented as string s, return the nesting depth of s.

// MaxDepth - calculates the maximum nesting depth of string s using custom filter, map and reduce HOFs
func MaxDepth(s string) int {
	if s == "" {
		return 0
	}

	chars := strings.Split(s, "")
	filtered := strFilter(chars, func(s string) bool {
		return s == "(" || s == ")"
	})

	mapped := strMap(filtered, func(s string) int {
		if s == "(" {
			return 1
		}

		return -1
	})

	scanned := reduceHOF(mapped, func(acc []interface{}, nxt int) []interface{} {
		nxt = acc[1].(int) + nxt
		acc[0] = append(acc[0].([]int), nxt)
		return []interface{}{acc[0], nxt}
	})

	sort.Ints(scanned[0].([]int))

	result := scanned[0]

	return result.([]int)[len(result.([]int))-1]
}

func strFilter(slice []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range slice {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func strMap(slice []string, f func(string) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func reduceHOF(arr []int, fn func([]interface{}, int) []interface{}) []interface{} {
	acc := []interface{}{[]int{}, 0}

	for _, v := range arr {
		acc = fn(acc, v)
	}

	return acc
}
