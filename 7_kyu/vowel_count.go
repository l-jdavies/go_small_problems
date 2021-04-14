/*
https://www.codewars.com/kata/54ff3102c1bad923760001f3

Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.
*/

package main

import "fmt"

func GetCount(str string) (count int) {
	count = 0

	for _, letter := range str {
		if IsVowel(string(letter)) {
			count++
		}
	}
	return
}

func IsVowel(l string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, v := range vowels {
		if l == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(GetCount("abracadabra") == 5)
}

/* These two solutions are much simplier:

Sol 1
func GetCount(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}

Sol 2
import "strings"

func GetCount(strn string) int {
  count := 0

  vowels := []string{"a", "e", "i", "o", "u"}

  for _, vowel := range vowels {
    count += strings.Count(strn, vowel)
  }

  return count
}
*/
