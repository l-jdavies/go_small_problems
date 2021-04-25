/* Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CountBits(n uint) int {
	bin := strconv.FormatInt(int64(n), 2)
	return strings.Count(bin, "1")
}

func main() {
	fmt.Println(CountBits(1234))
}

/* Alternative solution:
import "math/bits"

func CountBits(n uint) int {
  return bits.OnesCount(n)
} 
/*
