/*
The new "Avengers" movie has just been released! There are a lot of people at the cinema box office standing in a huge line. Each of them has a single 100, 50 or 25 dollar bill. An "Avengers" ticket costs 25 dollars.

Vasya is currently working as a clerk. He wants to sell a ticket to every single person in this line.

Can Vasya sell a ticket to every person and give change if he initially has no money and sells the tickets strictly in the order people queue?

Return YES, if Vasya can sell a ticket to every person and give change with the bills he has at hand at that moment. Otherwise return NO.

Examples:
Tickets([]int{25, 25, 50}) // => YES
Tickets([]int{25, 100}) // => NO. Vasya will not have enough money to give change to 100 dollars
Tickets([]int{25, 25, 50, 50, 100}) // => NO. Vasya will not have the right bills to give 75 dollars of chang
*/

package main

import "fmt"

func Tickets(peopleInLine []int) string {
	twentyFive := 0
	fifty := 0

	for _, i := range peopleInLine {
		switch i {
		case 25:
			twentyFive += 1
		case 50:
			fifty += 1
			twentyFive -= 1
		case 100:
			if fifty > 0 {
				fifty -= 1
				twentyFive -= 1
			} else {
				twentyFive -= 3
			}
		}

		if twentyFive < 0 || fifty < 0 {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	//fmt.Println(Tickets([]int{25, 25, 50, 50, 100}))
	//fmt.Println(Tickets([]int{25, 25, 50}))
	fmt.Println(Tickets([]int{25, 25, 50, 100}) == "YES")
	fmt.Println(Tickets([]int{25, 25, 25, 25, 50, 100, 50}) == "YES")
	fmt.Println(Tickets([]int{25, 25, 25, 100}) == "YES")
	fmt.Println(Tickets([]int{25, 50, 25, 100}) == "YES")
}
