/* Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').

Examples:

Solution("abc") //should return ["ab", "c_"]
Solution("abcdef") //should return ["ab", "cd", "ef"] */

package main

import "fmt"

func Solution(str string) []string {
	if len(str)%2 == 0 {
		return splitStr(str)
	} else {
		str += "_"
		return splitStr(str)
	}
}

func splitStr(str string) []string {
	if len(str) == 0 {
		return []string{}
	}
	arr := []string{}
	arr = append(arr, str[:2])
	arr = append(arr, splitStr(str[2:])...)
	return arr
}

func main() {
	fmt.Println(Solution("abc"))
	fmt.Println(Solution("abcdef"))
}
