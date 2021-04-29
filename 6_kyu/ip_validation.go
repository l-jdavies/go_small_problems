/* Write an algorithm that will identify valid IPv4 addresses in dot-decimal format. IPs should be considered valid if they consist of four octets, with values between 0 and 255, inclusive.

Input to the function is guaranteed to be a single string.

Examples
Valid inputs:

1.2.3.4
123.45.67.89
Invalid inputs:

1.2.3
1.2.3.4.5
123.456.78.90
123.045.067.089
*/

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Is_valid_ip(ip string) bool {
	if len(ip) == 0 {
		return false
	}
	a := regexp.MustCompile(`\.`).Split(ip, -1)

	alphWhite := regexp.MustCompile(`[[:alpha:]]|\s`) // match alphabetic and whitespace

	for _, v := range a {
		num, _ := strconv.Atoi(v)
		if num < 0 || num > 255 ||
			string(v[0]) == "0" && len(v) > 1 ||
			alphWhite.Match([]byte(v)) {
			return false
		}
	}

	return len(a) == 4
}

func main() {
	fmt.Println(Is_valid_ip("123.45.67.89"))
}

/* There's a Parse.IP function!:
import "net"
func Is_valid_ip(ip string) bool {
  res := net.ParseIP(ip)
  if res == nil {
    return false
  }
  return true
}
*/
