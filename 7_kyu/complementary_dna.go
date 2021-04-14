/*
https://www.codewars.com/kata/554e4a2f232cdd87d9000038/train/go

Deoxyribonucleic acid (DNA) is a chemical found in the nucleus of cells and carries the "instructions" for the development and functioning of living organisms.

If you want to know more http://en.wikipedia.org/wiki/DNA

In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". You have function with one side of the DNA (string, except for Haskell); you need to get the other complementary side. DNA strand is never empty or there is no DNA at all (again, except for Haskell).

More similar exercise are found here http://rosalind.info/problems/list-view/ (source)

DNA_strand ("ATTGC") # return "TAACG"

DNA_strand ("GTAT") # return "CATA"
*/

package main

import "fmt"

func DNAStrand(dna string) string {
	comp := ""

	for _, nucleo := range dna {
		switch string(nucleo) {
		case "A":
			comp = comp + "T"
		case "T":
			comp = comp + "A"
		case "C":
			comp = comp + "G"
		case "G":
			comp = comp + "C"
		}
	}
	return comp
}

func main() {
	fmt.Println(DNAStrand("AAAA") == "TTTT")
	fmt.Println(DNAStrand("ATTGC") == "TAACG")
	fmt.Println(DNAStrand("GTAT") == "CATA")
}

// note, could have used the strings.Replace method
