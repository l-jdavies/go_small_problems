/*
Insertion Sort consists of the following steps:
1. In the first pass-through, we temporarily remove the value at index 1 (the second cell) and store it in a temporary variable. This will leave a gap at that index, since it contains no value

In subsequent pass-throughs, we remove the values at the subsequent indexes.

2. We then begin a shifting phase, where we take each value to the left of the gap and compare it to the value in the temporary variable

If the value to the left of the gap is greater than the temporary variable, we shift that value to the right.
As we shift values to the right, inherently the gap moves leftward. As soon as we encounter a value that is lower than the temporarily removed value, or we reach the lft end of the array, this shifting phase is over.

3. We then insert the temporarily removed value into the current gap

4. Steps 1 through 3 represent a single pass-through. We repeat these pass- throughs until the pass-through begins at the final index of the array. By then, the array will have been fully sorted.

Like the bubble sort and selection sort algorithms, insertion sort has a time efficiency of O(N2). However, which algorithm is the fasted depends on the scenario. In the worst case scenario (array is in descending order and all elements need to be shifted), insertion sort is O(N2). In the average scenario, it's O(N2 / 2) and in the best case scenario (no elements need to be shifted), it's O(N). In contrast, the selection sort takes O(N/2) in all scenarios - each pass compares every value to the right of the chosen index, whatever happens.
*/

package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		position := i - 1

		for position >= 0 {
			if arr[position] > temp {
				arr[position+1] = arr[position]
				position--
			} else {
				break
			}
		}
		arr[position+1] = temp
	}
}

func main() {
	a := []int{1, 3, 11, 2, 8, 9}
	fmt.Println(a)
	insertionSort(a)
	fmt.Println(a)
}
