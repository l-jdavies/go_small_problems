/*
https://www.codewars.com/kata/5715eaedb436cf5606000381/train/go

You get an array of numbers, return the sum of all of the positives ones.

Example [1,-4,7,12] => 1 + 7 + 12 = 20

Note: if there is nothing to sum, the sum is default to 0.
*/
package main

import "fmt"

func PositiveSum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		if num > 0 {
			sum = sum + num
		}
	}
	return sum
}

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}) == 15)
	fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}) == 13)
	fmt.Println(PositiveSum([]int{}) == 0)
	fmt.Println(PositiveSum([]int{-1, -2, -3, -4, -5}) == 0)
	fmt.Println(PositiveSum([]int{-1, 2, 3, 4, -5}) == 9)
}
