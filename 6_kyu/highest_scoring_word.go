/* Given a string of words, you need to find the highest scoring word.

Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.

You need to return the highest scoring word as a string.

If two words score the same, return the word that appears earliest in the original string.

All letters will be lowercase and all inputs will be valid. */
package main

import (
	"fmt"
	"strings"
)

func High(s string) string {
	wordMap := getScoreMap(strings.Split(s, " "))
	fmt.Println(wordMap)

	highestVal := 0
	highestStr := ""

	for k, v := range wordMap {
		if v > highestVal {
			highestStr = k
			highestVal = v
		}
	}
	return highestStr
}

func getScoreMap(str []string) map[string]int {
	wordScores := map[string]int{}

	for _, v := range str {
		wordScores[string(v)] = getWordScore(string(v))
	}

	return wordScores
}

func getWordScore(word string) int {
	alph := "abcdefghijklmnopqrstuvwxyz"
	alphMap := map[string]int{}

	for i, v := range alph {
		alphMap[string(v)] = i + 1
	}

	score := 0
	for _, v := range word {
		score += alphMap[string(v)]
	}
	return score
}

func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
	fmt.Println(High("ricew yazslvy pichag bsdfnoc depedt kms ogu wfbie cgrzyvfeuo currpbqjtw"))

}
