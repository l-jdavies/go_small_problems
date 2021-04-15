/*
https://www.codewars.com/kata/566fc12495810954b1000030

Take an integer n (n >= 0) and a digit d (0 <= d <= 9) as an integer. Square all numbers k (0 <= k <= n) between 0 and n. Count the numbers of digits d used in the writing of all the k**2. Call nb_dig (or nbDig or ...) the function taking n and d as parameters and returning this count.

Examples:
n = 10, d = 1, the k*k are 0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
We are using the digit 1 in 1, 16, 81, 100. The total count is then 4.

nb_dig(25, 1):
the numbers of interest are
1, 4, 9, 10, 11, 12, 13, 14, 19, 21 which squared are 1, 16, 81, 100, 121, 144, 169, 196, 361, 441
so there are 11 digits `1` for the squares of numbers between 0 and 25.
Note that 121 has twice the digit 1.
*/

package main

import (
	"fmt"
	"strings"
)

func NbDig(n int, d int) int {
	squares := getSquares(n)
	strInt := convertIntToStr(squares)

	count := countDigit(strInt, d)
	return count
}

func countDigit(s []string, d int) int {
	count := 0

	for _, st := range s {
		count = count + strings.Count(st, fmt.Sprint(d))
	}
	return count
}

func convertIntToStr(n []int) []string {
	str := []string{}

	for _, num := range n {
		str = append(str, fmt.Sprint(num))
	}
	return str
}

func getSquares(n int) []int {
	s := []int{}

	for i := 0; i <= n; i++ {
		s = append(s, i*i)
	}
	return s
}

func main() {
	fmt.Println(NbDig(10, 1) == 4)
	fmt.Println(NbDig(25, 1) == 11)
	fmt.Println(NbDig(550, 5) == 213)
}

/* alternative solution:
import (
  "strings"
  "strconv"
)

func NbDig(n int, d int) (out int) {
  for i := 0; i <= n; i++ {
    out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
  }

  return
}

strconv.Itoa() converts string to int64. This replaces my `convertIntToStr` function
*/
