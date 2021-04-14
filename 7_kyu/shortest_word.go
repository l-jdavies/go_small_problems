/*
https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9

Simple, given a string of words, return the length of the shortest word(s).

String will never be empty and you do not need to account for different data types.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func FindShort(s string) int {
	sliceStr := strings.Split(s, " ")
	sort.Sort(byLength(sliceStr))
	return len(sliceStr[0])
}

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps") == 3)
	fmt.Println(FindShort("turns out random test cases are easier than writing out basic ones") == 3)
}
