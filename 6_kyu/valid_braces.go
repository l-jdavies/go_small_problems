/* Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.

This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!

All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

What is considered Valid?
A string of braces is considered valid if all braces are matched with the correct brace.

Examples
"(){}[]"   =>  True
"([{}])"   =>  True
"(}"       =>  False
"[(])"     =>  False
"[({})](]" =>  False
*/

package main

import (
	"fmt"
	"regexp"
)

func ValidBraces(str string) bool {
	closing := map[string]string{"(": ")", "{": "}", "[": "]"}
	braceCollection := []string{}

	for _, v := range str {
		open, _ := regexp.MatchString(`\(|\{|\[`, string(v))
		closed, _ := regexp.MatchString(`\)|\}|\]`, string(v))

		if open {
			braceCollection = append(braceCollection, string(v))
		} else if closed {
			if len(braceCollection) == 0 {
				return false
			}
			lastBrace := braceCollection[len(braceCollection)-1]
			braceCollection = braceCollection[:len(braceCollection)-1]
			if closing[lastBrace] != string(v) {
				return false
			}
		}
	}

	if len(braceCollection) > 0 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(ValidBraces("(){}[]"))
	fmt.Println(ValidBraces("([{}])"))
	fmt.Println(ValidBraces("(})"))
}

/* Cleaner solution:
func ValidBraces(str string) bool {
  stack := make([]rune, 0)
  for _, c := range str {
    switch c {
    case '(':
      stack = append(stack, ')')
    case '[':
      stack = append(stack, ']')
    case '{':
      stack = append(stack, '}')
    case ')', ']', '}':
      if len(stack) == 0 {
        return false
      }
      if stack[len(stack)-1] != c {
        return false
      }
      stack = stack[:len(stack)-1]
    }
  }
  return len(stack) == 0
}
*/
