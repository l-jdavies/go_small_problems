/*
A binary search can only be used on an ordered array.

It essentially selects a midpoint in the array and determines if the target value is higher or lower than the midpoint value. Based on that result, half of the array will be eliminated from the search and the process begins again.

The time complexity of a binary search is O(log N) - doubling the number of elements in the ordered array will only increase the number of algorithm steps by 1.
*/

package main

import "fmt"

func binarySearch(arr []int, target int) int {
	lower := 0
	upper := len(arr) - 1

	for lower <= upper {
		midpoint := (lower + upper) / 2
		valueAtMidpoint := arr[midpoint]

		if valueAtMidpoint == target {
			return midpoint
		} else if target < valueAtMidpoint {
			upper = midpoint - 1
		} else if target > valueAtMidpoint {
			lower = midpoint + 1
		}

	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6}, 4))
	fmt.Println(binarySearch([]int{1, 3, 4, 5, 6, 7}, 2))
}
