// A function that determines if a slice contains duplicate values is commonly solved with a nested loop, which takes quadratic time. However, this can be solved with a linear solution:

package main

import "fmt"

func hasDuplicates(arr []int) bool {
	hsh := map[int]bool{}

	for _, v := range arr {
		if hsh[v] == true {
			return true
		} else {
			hsh[v] = true
		}
	}
	return false
}

func main() {
	fmt.Println(hasDuplicates([]int{1, 2, 3, 1}))
	fmt.Println(hasDuplicates([]int{1, 2, 3, 4}))
}
