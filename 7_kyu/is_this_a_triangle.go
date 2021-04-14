/*
https://www.codewars.com/kata/56606694ec01347ce800001b

Implement a method that accepts 3 integer values a, b, c. The method should return true if a triangle can be built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).
*/

package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsTriangle(5, 1, 2) == false)
	fmt.Println(IsTriangle(1, 2, 5) == false)
	fmt.Println(IsTriangle(2, 5, 1) == false)
	fmt.Println(IsTriangle(4, 2, 3) == true)
	fmt.Println(IsTriangle(5, 1, 5) == true)
}
