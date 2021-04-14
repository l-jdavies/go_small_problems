/*
codewars.com/kata/55fd2d567d94ac3bc9000064/train/go

Given the triangle of consecutive odd numbers:

             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29
...
Calculate the row sums of this triangle from the row index (starting at index 1) e.g.:

rowSumOddNumbers(1); // 1
rowSumOddNumbers(2); // 3 + 5 = 8
*/

package main

import "fmt"

func RowSumOddNumbers(n int) int {
	triangle := createTriangle(n)
	lastRow := triangle[len(triangle)-1]
	sum := sumRow(lastRow)
	return sum
}

func sumRow(s []int) int {
	sum := 0
	for _, element := range s {
		sum = sum + element
	}
	return sum
}

func createTriangle(n int) [][]int {
	outer := [][]int{}
	num := 1

	for rowNum := 0; len(outer) <= n; rowNum++ {
		inner := []int{}
		for len(inner) < rowNum {
			inner = append(inner, num)
			num += 2
		}
		outer = append(outer, inner)
	}
	return outer
}

func main() {
	fmt.Println(RowSumOddNumbers(1) == 1)
	fmt.Println(RowSumOddNumbers(2) == 8)
}
