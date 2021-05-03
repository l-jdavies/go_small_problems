/*
Implement the following in golang with an improved time efficiency (currently O(N2))

The following function checks whether an array of numbers contains a pair of two numbers that add up to 10.

function twoSum(array) {
  for (let i = 0; i < array.length; i++) {
    for (let j = 0; j < array.length; j++) {
      if (i !== j && array[i] + array[j] === 10) {
      return true;
      }
    }
  }
  return false;
}
*/

package main

import "fmt"

func twoSum(arr []int) bool {
	hsh := map[int]bool{}

	for _, v := range arr {
		hsh[10-v] = true
	}

	for _, v := range arr {
		if hsh[v] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(twoSum([]int{3, 1, 2, 11, 8}))
	fmt.Println(twoSum([]int{3, 1, 11, 8}))
}

/*
hsh = {
  7: true,
  9: true,
  8: true,
  -1: true,
  2: true
}
*/
