/*
https://www.codewars.com/kata/56b7f2f3f18876033f000307

In this Kata, your function receives an array of integers as input. Your task is to determine whether the numbers are in ascending order. An array is said to be in ascending order if there are no two adjacent integers where the left integer exceeds the right integer in value.

For the purposes of this Kata, you may assume that all inputs are valid, i.e. non-empty arrays containing only integers.

Note that an array of 1 integer is automatically considered to be sorted in ascending order since all (zero) adjacent pairs of integers satisfy the condition that the left integer does not exceed the right integer in value. An empty list is considered a degenerate case and therefore will not be tested in this Kata - feel free to raise an Issue if you see such a list being tested.

For example:

InAscOrder([]int{1, 2, 4, 7, 19}) // returns True
InAscOrder([]int{1, 2, 3, 4, 5}) // returns True
InAscOrder([]int{1, 6, 10, 18, 2, 4, 20}) // returns False
InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}) // returns False because the numbers are in DESCENDING order
*/

package main

import (
	"fmt"
	"sort"
)

func InAscOrder(number []int) bool {
	return sort.IntsAreSorted(number)
}

func main() {
	fmt.Println(InAscOrder([]int{1, 2, 4, 7, 19}) == true)
	fmt.Println(InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}) == false)
}
