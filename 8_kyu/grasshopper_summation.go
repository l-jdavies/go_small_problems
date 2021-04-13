/*
https://www.codewars.com/kata/55d24f55d7dd296eb9000030/train/go

Write a program that finds the summation of every number from 1 to num. The number will always be a positive integer greater than 0.

For example:

summation(2) -> 3
1 + 2

summation(8) -> 36
1 + 2 + 3 + 4 + 5 + 6 + 7 + 8
*/

package main

import "fmt"

func Summation(n int) int {
	sum := 0

	for current := 1; current <= n; current++ {
		sum = sum + current
	}
	return sum
}

func main() {
	fmt.Println(Summation(1) == 1)
	fmt.Println(Summation(8) == 36)
	fmt.Println(Summation(22) == 253)
	fmt.Println(Summation(100) == 5050)
	fmt.Println(Summation(213) == 22791)
}
