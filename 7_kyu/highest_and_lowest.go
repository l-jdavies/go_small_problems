/*
https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go

In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

Example:

highAndLow("1 2 3 4 5");  // return "5 1"
highAndLow("1 2 -3 4 5"); // return "5 -3"
highAndLow("1 9 3 4 -5"); // return "9 -5"
Notes:

All numbers are valid Int32, no need to validate them.
There will always be at least one number in the input string.
Output string must be two numbers separated by a single space, and highest number is first.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	sliceStr := strings.Split(in, " ")
	sliceInt := []int{}

	for _, char := range sliceStr {
		num, _ := strconv.ParseInt(string(char), 0, 64)
		sliceInt = append(sliceInt, int(num))
	}

	sort.Ints(sliceInt)
	result := fmt.Sprint(sliceInt[len(sliceInt)-1])
	result = result + " " + fmt.Sprint(sliceInt[0])
	return result
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5") == "5 1")
	fmt.Println(HighAndLow("1 2 -3 4 5") == "5 -3")
	fmt.Println(HighAndLow("1 9 3 4 -5") == "9 -5")
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4") == "42 -9")
}
