/*
https://www.codewars.com/kata/5aba780a6a176b029800041c

Task
Given a Divisor and a Bound , Find the largest integer N , Such That ,

Conditions :
N is divisible by divisor

N is less than or equal to bound

N is greater than 0.

Notes
The parameters (divisor, bound) passed to the function are only positive values .
It's guaranteed that a divisor is Found .
*/

package main

import "fmt"

func MaxMultiple(d, b int) int {
	for n := b; n > 0; n-- {
		if n%d == 0 {
			return n
		}
	}
	return 0
}

func main() {
	fmt.Println(MaxMultiple(2, 7) == 6)
	fmt.Println(MaxMultiple(3, 10) == 9)
}
