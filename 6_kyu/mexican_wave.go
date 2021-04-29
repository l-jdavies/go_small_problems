/* In this simple Kata your task is to create a function that turns a string into a Mexican Wave. You will be passed a string and you must return that string in an array where an uppercase letter is a person standing up.
Rules
 1.  The input string will always be lower case but maybe empty.

 2.  If the character in the string is whitespace then pass over it as if it was an empty seat
Example
wave("hello") => []string{"Hello", "hEllo", "heLlo", "helLo", "hellO"} */

package main

import (
	"fmt"
	"strings"
)

func wave(words string) []string {
	noSpace := strings.Replace(words, " ", "", -1)
	r := []string{}
	upper := 0

	for len(r) < len(noSpace) {
		if string(words[upper]) != " " {
			temp := words[:upper] + strings.ToUpper(string(words[upper])) + words[upper+1:]
			r = append(r, temp)
			upper++
		} else {
			upper++
		}
	}
	return r
}

func main() {
	fmt.Println(wave(" x yz"))
}
