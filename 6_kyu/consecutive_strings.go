/*
You are given an array(list) strarr of strings and an integer k. Your task is to return the first longest string consisting of k consecutive strings taken in the array.

Examples:
strarr = ["tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"], k = 2

Concatenate the consecutive strings of strarr by 2, we get:

treefoling   (length 10)  concatenation of strarr[0] and strarr[1]
folingtrashy ("      12)  concatenation of strarr[1] and strarr[2]
trashyblue   ("      10)  concatenation of strarr[2] and strarr[3]
blueabcdef   ("      10)  concatenation of strarr[3] and strarr[4]
abcdefuvwxyz ("      12)  concatenation of strarr[4] and strarr[5]

Two strings are the longest: "folingtrashy" and "abcdefuvwxyz".
The first that came is "folingtrashy" so
longest_consec(strarr, 2) should return "folingtrashy".

In the same way:
longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"
*/

package main

import (
	"fmt"
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	if len(strarr) == 0 ||
		k > len(strarr) ||
		k <= 0 {
		return ""
	}

	consec := getConsec(strarr, k)

	return findLongest(consec)
}

func findLongest(arr []string) string {
	longest := ""

	for _, v := range arr {
		if len(string(v)) > len(longest) {
			longest = string(v)
		}
	}
	return longest
}

func getConsec(arr []string, k int) []string {
	a := []string{}
	idx1 := 0
	idx2 := k

	for idx2 <= len(arr) {
		a = append(a, strings.Join(arr[idx1:idx2], ""))
		idx1++
		idx2++
	}
	return a
}

func main() {
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"}, 2))
}

/* More concise solution:
import ("strings")

func LongestConsec(strarr []string, k int) string {
  var buffer string
  var largest string

  for i := 0; i <= len(strarr)-k; i++ {
    buffer = strings.Join(strarr[i : i+k][:], "")
    if len(buffer) > len(largest) {
      largest = buffer
    }
  }
  return largest
}
*/
