/*
Complete the method/function so that it converts dash/underscore delimited words into camel casing. The first word within the output should be capitalized only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case).

Examples
"the-stealth-warrior" gets converted to "theStealthWarrior"
"The_Stealth_Warrior" gets converted to "TheStealthWarrior"
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	re := regexp.MustCompile("[-_]")
	e := re.ReplaceAllString(s, " ")
	slice := strings.Split(e, " ")

	for i, w := range slice {
		if i == 0 {
			slice[i] = w
		} else {
			slice[i] = strings.Title(w)
		}
	}
	return strings.Join(slice, "")
}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
}

/* Cleaner solution chaining the Split function onto MustCompile:
func ToCamelCase(s string) string {
  words := regexp.MustCompile("-|_").Split(s, -1)

  for i, w := range words[1:] {
    words[i+1] = strings.Title(w)
  }

  return strings.Join(words, "")
}
*/
