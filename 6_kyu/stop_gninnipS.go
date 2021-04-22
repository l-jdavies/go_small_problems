/*
Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed (like the name of this kata).

Strings passed in will consist of only letters and spaces.
Spaces will be included only when more than one word is present.
Examples:

spinWords("Hey fellow warriors") => "Hey wollef sroirraw"
spinWords("This is a test") => "This is a test"
spinWords("This is another test") => "This is rehtona test"
*/

package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	s := strings.Split(str, " ")
	r := []string{}

	for _, word := range s {
		if len(word) >= 5 {
			rev := reverseString(word)
			r = append(r, rev)
		} else {
			r = append(r, word)
		}
	}
	return strings.Join(r, " ")
}

func reverseString(s string) string {
	result := ""

	for _, l := range s {
		result = string(l) + result
	}
	return result
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))
}
