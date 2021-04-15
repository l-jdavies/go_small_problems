/*
https://www.codewars.com/kata/5663f5305102699bad000056

You are given two arrays a1 and a2 of strings. Each string is composed with letters from a to z. Let x be any string in the first array and y be any string in the second array.

Find max(abs(length(x) − length(y)))

If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).

Example:
a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
mxdiflg(a1, a2) --> 13

*/

package main

import "fmt"

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	longestA := findLongest(a1)
	shortestA := findShortest(a1)
	longestB := findLongest(a2)
	shortestB := findShortest(a2)

	max := findMaxDiff(longestA, longestB, shortestA, shortestB)
	return max
}

func findMaxDiff(longA, longB, shortA, shortB int) int {
	a := longA - shortB
	b := longB - shortA

	if a > b {
		return a
	} else {
		return b
	}
}

func findLongest(a []string) int {
	longest := ""
	max := -1

	for _, l := range a {
		if len(l) < max {
			continue
		} else {
			longest = l
			max = len(l)
		}
	}
	return len(longest)
}

func findShortest(a []string) int {
	shortest := ""
	min := 100

	for _, l := range a {
		if len(l) < min {
			shortest = l
			min = len(l)
		} else {
			continue
		}
	}
	return len(shortest)
}

func main() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	fmt.Println(MxDifLg(a1, a2) == 13)
}
