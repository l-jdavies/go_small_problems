/*
The two oldest ages function/method needs to be completed. It should take an array of numbers as its argument and return the two highest numbers within the array. The returned value should be an array in the format [second oldest age, oldest age].

The order of the numbers passed in could be any order. The array will always include at least 2 items. If there are two or more oldest age, then return both of them in array format.

For example:

TwoOldestAges([]int{1, 5, 87, 45, 8, 8}) // should return [2]int{45, 87}
*/

package main

import (
	"fmt"
	"sort"
)

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	y := ages[len(ages)-2]
	o := ages[len(ages)-1]
	return [2]int{y, o}
}

func main() {
	fmt.Println(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}))
}

/* alternative solution
func TwoOldestAges(ages []int) [2]int {
  a, b := 0, 0
  for _, v := range ages {
    if v > b {
      a, b = b, v
    } else if v > a {
      a = v
    }
  }
  return [2]int{a, b}
}
*/
