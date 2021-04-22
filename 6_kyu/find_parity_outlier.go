/*
You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N. Write a method that takes the array as an argument and returns this "outlier" N.

Examples
[2, 4, 0, 100, 4, 11, 2602, 36]
Should return: 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21]
Should return: 160 (the only even number)
*/

package main

import "fmt"

func FindOutlier(integers []int) (outlier int) {
	maj := OddOrEven(integers)

	switch maj {
	case "odd":
		outlier = findEven(integers)
	case "even":
		outlier = findOdd(integers)
	}
	return
}

func findOdd(ints []int) (odd int) {
	for _, n := range ints {
		if n%2 != 0 {
			odd = n
			break
		}
	}
	return
}

func findEven(ints []int) (even int) {
	for _, n := range ints {
		if n%2 == 0 {
			even = n
			break
		}
	}
	return
}

func OddOrEven(ints []int) string {
	countOdd := 0

	for _, n := range ints {
		if n%2 != 0 {
			countOdd++
		}
	}

	if countOdd > 1 {
		return "odd"
	} else {
		return "even"
	}
}

func main() {
	fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
}
