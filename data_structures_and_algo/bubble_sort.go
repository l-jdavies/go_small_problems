/*
Bubble Sort is a basic sorting algorithm and follows these steps:
1. Point to two consecutive values in the array. (Initially, we start by pointing
to the array’s first two values.) Compare the first item with the second one

2. If the two items are out of order (in other words, the left value is greater than the right value), swap them (if they already happen to be in the cor- rect order, do nothing for this step)

3. Move the “pointers” one cell to the right

4. Repeat Steps 1 through 3 until we reach the end of the array, or if we reach the values that have already been sorted. (This will make more sense in the walk-through that follows.) At this point, we have completed our first pass-through of the array. That is, we “passed through” the array by pointing to each of its values until we reached the end.

5. We then move the two pointers back to the first two values of the array, and execute another pass-through of the array by running Steps 1 through 4 again. We keep on executing these pass-throughs until we have a pass- through in which we did not perform any swaps. When this happens, it means our array is fully sorted and our work is done.

Time efficiency is O(N2), which is quadratic time (lot less efficient than O(N))
*/

package main

import "fmt"

func bubbleSort(l []int) {
	rightIdx := len(l) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < rightIdx; i++ {
			if l[i] > l[i+1] && i+1 < len(l) {
				l[i], l[i+1] = l[i+1], l[i]
				sorted = false
			}
		}
		rightIdx -= 1
	}
}

func main() {
	arr := []int{65, 55, 45, 35, 25, 15, 10}
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}
