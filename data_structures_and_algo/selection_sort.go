/*
The steps of Selection Sort are as follows:
1. We check each cell of the array from left to right to determine which value is least. As we move from cell to cell, we keep track of the lowest value we’ve encountered so far. (We’ll do this by storing its index in a variable.) If we encounter a cell that contains a value that is even lower than the one in our variable, we replace it so that the variable now points to the new index.

 2. Once we’ve determined which index contains the lowest value, we swap its value with the value we began the pass-through with. This would be index 0 in the first pass-through, index 1 in the second pass-through, and so on. The diagram here illustrates making the swap of the first pass-through.

3. Each pass-through consists of Steps 1 and 2.

We repeat the pass-throughs until we reach a pass-through that would start at the end of the array. By this point, the array will have been fully sorted.

In terms of time efficiency, selection sort algorithm makes few comparisons and performs fewer swaps than the bubble sort algorithm, so technically the time efficiency is O(N2 / 2); however because constants are ignored, the selection sort algorithm also has a time efficiency of O(N2)
*/

package main

import "fmt"

func selectionSort(list []int) {
	for i, _ := range list {
		lowestValIdx := i
		for idx := i + 1; idx < len(list); idx++ {
			if list[idx] < list[lowestValIdx] {
				lowestValIdx = idx
			}
		}

		if lowestValIdx != i {
			list[i], list[lowestValIdx] = list[lowestValIdx], list[i]
		}
	}
}

func main() {
	arr := []int{2, 6, 1, 3}
	fmt.Println(arr)
	selectionSort(arr)
	fmt.Println(arr)

	arr2 := []int{4, 2, 7, 1, 3}
	fmt.Println(arr2)
	selectionSort(arr2)
	fmt.Println(arr2)
}
