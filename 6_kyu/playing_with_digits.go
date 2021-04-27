/*Some numbers have funny properties. For example:

89 --> 8¹ + 9² = 89 * 1

695 --> 6² + 9³ + 5⁴= 1390 = 695 * 2

46288 --> 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51

Given a positive integer n written as abcd... (a, b, c, d... being digits) and a positive integer p

we want to find a positive integer k, if it exists, such as the sum of the digits of n taken to the successive powers of p is equal to k * n.
In other words:

Is there an integer k such as : (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k

If it is the case we will return k, if not return -1.

Note: n and p will always be given as strictly positive integers.

*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func digPow(n, p int) int {
	str := fmt.Sprint(n)
	dig := strings.Split(str, "")
	sum := 0.0

	for _, i := range dig {
		x, _ := strconv.ParseFloat(string(i), 64)
		sum += math.Pow(x, float64(p))
		p++
	}

	k := -1

	for mult := 1; mult <= n; mult++ {
		if int(sum)%n == 0 {
			k = int(sum) / n
			break
		}
	}
	return k
}

func main() {
	fmt.Println(digPow(695, 2) == 2)
	fmt.Println(digPow(46288, 3) == 51)
	fmt.Println(digPow(92, 1) == -1)
}
