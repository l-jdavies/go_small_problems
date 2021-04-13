/*
https://www.codewars.com/kata/50654ddff44f800200000004/train/go

This code does not execute properly. Try to figure out why.

package multiply

func Multiply(a, b int) int {
  a * b
}

answer = needed a return statement
*/
package main

import "fmt"

func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Multiply(1, 1) == 1)
	fmt.Println(Multiply(1, 0) == 0)
	fmt.Println(Multiply(2, 5) == 10)
	fmt.Println(Multiply(5, 10) == 50)
}
