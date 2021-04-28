/* There is an array with some numbers. All numbers are equal except for one. Try to find it!

findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
It’s guaranteed that array contains at least 3 numbers.

The tests contain some very huge arrays, so think about performance. */
package main

import "fmt"

func main() {
	fmt.Println(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}))
}

func FindUniq(arr []float32) float32 {
	hsh := map[float32]int{}
	var uniq float32

	for _, v := range arr {
		hsh[v]++
	}

	for k, v := range hsh {
		if v == 1 {
			uniq = k
			break
		}
	}
	return uniq
}
