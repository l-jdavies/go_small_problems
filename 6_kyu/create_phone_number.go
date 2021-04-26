/*
Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.

Example:
CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
*/

package main

import (
	"fmt"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	str := convertToString(numbers)
	num := "(" + strings.Join(str[:3], "") + ") "
	num += strings.Join(str[3:6], "") + "-"
	num += strings.Join(str[6:], "")
	return num
}

func convertToString(n [10]uint) []string {
	s := []string{}

	for _, i := range n {
		s = append(s, fmt.Sprint(i))
	}
	return s
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

/* Cleaner solution:
func ArrayToString(numbers interface{}) string {
  return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

func CreatePhoneNumber(numbers [10]uint) string {
  str := ArrayToString(numbers)
  return fmt.Sprintf("(%s) %s-%s", str[0:3], str[3:6], str[6:10])
}
*/
