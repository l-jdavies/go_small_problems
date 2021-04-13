/*
https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go

Complete the solution so that it reverses the string passed into it.

'world'  =>  'dlrow'
*/

package main

import "fmt"

func Solution(word string) (reversed string) {
	for _, l := range word {
		reversed = string(l) + reversed
	}
	return
}

func main() {
	fmt.Println(Solution("world") == "dlrow")
}
