/*
https://www.codewars.com/kata/583f158ea20cfcbeb400000a

Given two numbers and an arithmetic operator (the name of it, as a string), return the result of the two numbers having that operator used on them.

a and b will both be positive integers, and a will always be the first number in the operation, and b always the second.

The four operators are "add", "subtract", "divide", "multiply".

A few examples:

Arithmetic(5, 2, "add")      => returns 7
Arithmetic(5, 2, "subtract") => returns 3
Arithmetic(5, 2, "multiply") => returns 10
Arithmetic(5, 2, "divide")   => returns 2

Try to do it without using if statements!
*/

package main

import "fmt"

func Arithmetic(a int, b int, operator string) int {
	result := 0
	switch operator {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		result = a / b
	}
	return result
}

func main() {
	fmt.Println(Arithmetic(8, 2, "add") == 10)
	fmt.Println(Arithmetic(8, 2, "subtract") == 6)
	fmt.Println(Arithmetic(8, 2, "multiply") == 16)
}
