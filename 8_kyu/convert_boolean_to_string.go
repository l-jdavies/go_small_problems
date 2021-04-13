/*
https://www.codewars.com/kata/53369039d7ab3ac506000467/train/go

Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false

*/

package main

import "fmt"

func BoolToWord(word bool) (result string) {
	switch word {
	case true:
		result = "Yes"
	case false:
		result = "No"
	}
	return
}

func main() {
	fmt.Println(BoolToWord(true) == "Yes")
	fmt.Println(BoolToWord(false) == "No")
}

/* cleaner logic:
func BoolToWord(word bool) string {
    if word {
      return "Yes"
    }

    return "No"
}
*/
