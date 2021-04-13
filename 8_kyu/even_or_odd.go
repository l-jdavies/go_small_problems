//  Create a function (or write a script in Shell) that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.

package main

import "fmt"

func EvenOrOdd(number int) string {
	division := number % 2

	switch division {
	case 0:
		return "Even"
	default:
		return "Odd"
	}
}

func main() {
	fmt.Println(EvenOrOdd(2))
	fmt.Println(EvenOrOdd(3))
	fmt.Println(EvenOrOdd(4))
	fmt.Println(EvenOrOdd(10))
}
