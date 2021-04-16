/*
https://www.codewars.com/kata/59cfc000aeb2844d16000075

Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below. Index 0 will be considered even.

For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF']. See test cases for more examples.

The input will be a lowercase string with no spaces.

*/

package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string) []string {
	s := []string{}

	s = append(s, capitalizeEven(st))
	s = append(s, capitalizeOdd(st))
	return s
}

func capitalizeEven(st string) string {
	c := ""
	for i, l := range st {
		if i%2 == 0 {
			c = c + strings.ToUpper(string(l))
		} else {
			c = c + string(l)
		}
	}
	return c
}

func capitalizeOdd(st string) string {
	c := ""
	for i, l := range st {
		if i%2 != 0 {
			c = c + strings.ToUpper(string(l))
		} else {
			c = c + string(l)
		}
	}
	return c
}

func main() {
	fmt.Println(Capitalize("abcdef"))
}

/*
Alternative solution

func Capitalize(st string) []string {
  s1 := ""
  s2 := ""
  for i, c :=range st {
    if i%2==0 {
      s1 += strings.ToUpper(string(c))
      s2 += strings.ToLower(string(c))
    } else {
      s1 += strings.ToLower(string(c))
      s2 += strings.ToUpper(string(c))
    }
  }

  return []string{s1,s2}
}
*/
