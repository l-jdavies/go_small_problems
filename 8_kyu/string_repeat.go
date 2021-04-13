/*
https://www.codewars.com/kata/57a0e5c372292dd76d000d7e/train/go

Write a function called repeat_string which repeats the given string str exactly count times.

repeatStr(6, "I") // "IIIIII"
repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"
*/

package main

import (
	"fmt"
	s "strings"
)

func RepeatStr(rep int, value string) string {
	return s.Repeat(value, rep)
}

func main() {
	fmt.Println(RepeatStr(4, "a") == "aaaa")
	fmt.Println(RepeatStr(3, "hello") == "hellohellohello")
	fmt.Println(RepeatStr(2, "abc") == "abcabc")
}
