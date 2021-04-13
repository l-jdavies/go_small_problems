/*
https://www.codewars.com/kata/55685cd7ad70877c23000102/train/go

In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

Example:

MakeNegative(1)    # return -1
MakeNegative(-5)   # return -5
MakeNegative(0)    # return 0
Notes:

The number can be negative already, in which case no change is required.
Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.
*/

package main

import "fmt"

func MakeNegative(x int) int {
	if x <= 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	fmt.Println(MakeNegative(1) == -1)
	fmt.Println(MakeNegative(-5) == -5)
	fmt.Println(MakeNegative(0) == 0)
	fmt.Println(MakeNegative(42) == -42)
}
