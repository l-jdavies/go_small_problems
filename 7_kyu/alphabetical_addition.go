/*
Your task is to add up letters to one letter.

The function will be given a slice of letters (runes), each one being a letter to add. Return a rune.

Notes:
Letters will always be lowercase.
Letters can overflow (see second to last example of the description)
If no letters are given, the function should return 'z'
Examples:
AddLetters([]rune{'a', 'b', 'c'}) = 'f'
AddLetters([]rune{'a', 'b'}) = 'c'
AddLetters([]rune{'z'}) = 'z'
AddLetters([]rune{'z', 'a'}) = 'a'
AddLetters([]rune{'y', 'c', 'b'}) = 'd' // notice the letters overflowing
AddLetters([]rune{}) = 'z'
*/

package main

import (
	"fmt"
)

func AddLetters(letters []rune) rune {
	if len(letters) == 0 {
		return rune(122)
	}

	var sum rune = 0

	for _, r := range letters {
		sum = sum + rune((r - 96))
	}

	result := (sum-1)%26 + 97
	return rune(result)
}

func main() {
	fmt.Println(AddLetters([]rune{'x', 'w', 's', 'b', 'w', 'z'}))
}
