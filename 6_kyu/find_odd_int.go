/*
Given an array of integers, find the one that appears an odd number of times.

There will always be only one integer that appears an odd number of times.

*/
package main

import (
	"fmt"
)

func FindOdd(seq []int) int {
	counts := createMapCount(seq)
	odd := 0

	for k, v := range counts {
		if v%2 != 0 {
			return k
		}
	}
	return odd
}

func createMapCount(seq []int) map[int]int {
	count := map[int]int{}

	for _, n := range seq {
		if count[n] > 0 {
			count[n] += 1
		} else {
			count[n] = 1
		}
	}
	return count
}

func main() {
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
	fmt.Println(FindOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}))
}
