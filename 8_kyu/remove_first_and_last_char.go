/*
https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0/train/go

It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string. You're given one parameter, the original string. You don't have to worry with strings with less than two characters.
*/

package main

import (
	"fmt"
	s "strings"
)

func RemoveChar(word string) string {
	slice := s.Split(word, "")
	slice = slice[1:]
	slice = slice[:len(slice)-1]

	return s.Join(slice, "")
}

func main() {
	fmt.Println(RemoveChar("hello") == "ell")
	fmt.Println(RemoveChar("eloquent") == "loquen")
	fmt.Println(RemoveChar("country") == "ountr")
}

/* simplier syntax, using index position of string:
func RemoveChar(word string) string {
  return word[1:len(word)-1]
}
*/
