/* Given two arrays of strings a1 and a2 return a sorted array r in lexicographical order of the strings of a1 which are substrings of strings of a2.

#Example 1: a1 = ["arp", "live", "strong"]

a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

returns ["arp", "live", "strong"]

#Example 2: a1 = ["tarp", "mice", "bull"]

a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

returns []

Notes:
Arrays are written in "general" notation. See "Your Test Cases" for examples in your language.

In Shell bash a1 and a2 are strings. The return is a string where words are separated by commas.

Beware: r must be without duplicates.
Don't mutate the inputs. */

package main

import (
	"fmt"
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	r := []string{}

	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if strings.Contains(v2, v1) {
				r = append(r, v1)
				break
			}
		}
	}

	r = getUnique(r)
	sort.Strings(r)
	return r
}

func getUnique(arr []string) []string {
	hsh := map[string]int{}
	uniq := []string{}

	for _, v := range arr {
		if hsh[v] == 0 {
			uniq = append(uniq, v)
			hsh[v]++
		}
	}
	return uniq
}

func main() {
	var a1 = []string{"live", "arp", "strong"}
	var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(InArray(a1, a2))
}
