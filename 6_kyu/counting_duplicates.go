/*
Count the number of Duplicates
Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that occur more than once in the input string. The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.

Example
"abcde" -> 0 # no characters repeats more than once
"aabbcde" -> 2 # 'a' and 'b'
"aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
"indivisibility" -> 1 # 'i' occurs six times
"Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
"aA11" -> 2 # 'a' and '1'
"ABBA" -> 2 # 'A' and 'B' each occur twice
*/

package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s string) int {
	countMap := countCharacters(strings.ToLower(s))
	count := 0

	for _, v := range countMap {
		if v > 1 {
			count++
		}
	}
	return count
}

func countCharacters(s string) map[string]int {
	countMap := map[string]int{}

	for _, i := range s {
		countMap[string(i)]++
	}
	return countMap
}

func main() {
	fmt.Println(duplicate_count("abcde"))
	fmt.Println(duplicate_count("aabBcde"))
}
