/*
https://www.codewars.com/kata/56dec885c54a926dcd001095/train/go

Given a number, find its opposite:
1: -1
14: -14
-34: 34
*/

package main

import (
	"fmt"
	"math"
)

func Opposite(value float64) int {
	if value < 0 {
		return int(math.Abs(value))
	} else {
		return int(value) - int(2*value)
	}
}

func main() {
	fmt.Println(Opposite(1) == -1)
	fmt.Println(Opposite(14) == -14)
	fmt.Println(Opposite(-34) == 34)
}

/* best solution:
func Opposite(value int) int {
  return -value
	}
*/
