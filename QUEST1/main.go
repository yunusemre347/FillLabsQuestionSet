package main

import (
	"fmt"
	"sort"
	"strings"
)

// create a struct for later use. We will use this to give a counter for every member
type newWord struct {
	word       string
	counterOfA int
}

// create a slice for the final wrap-up
var finalWords = []string{}

func main() {
	var words = []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	//var words = []string{"aaaasd", "ab", "a", "adsf", "asfdfshgdfh", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	//this will be our slice of structs to see every words "a" counter.
	var newWords = make([]newWord, 0)
	//this slice is to sort the rest of the words according to their length
	var wordsWithoutA = make([]newWord, 0)
	//this loop is to check if they have "a" amongst the letters & assign a counter
	for _, word := range words {
		var counterOfA int = 0
		if strings.Contains(word, "a") {
			var letters = strings.Split(word, "")
			for _, letter := range letters {
				if letter == "a" {
					counterOfA++
				}
			}
			var data = newWord{
				word:       word,
				counterOfA: counterOfA,
			}
			newWords = append(newWords, data)
		} else {
			var data = newWord{
				word:       word,
				counterOfA: counterOfA,
			}
			wordsWithoutA = append(wordsWithoutA, data)
		}
	}
	//sort the words with "a" according to their length
	sort.Slice(newWords, func(i, j int) bool {

		return len(newWords[i].word) > len(newWords[j].word)
	})
	//sort the words with "a" according to their "a"count.
	sort.Slice(newWords, func(i, j int) bool {

		return newWords[i].counterOfA > newWords[j].counterOfA
	})
	//sort the words without an "a" according to their length
	sort.Slice(wordsWithoutA, func(i, j int) bool {

		return len(wordsWithoutA[i].word) > len(wordsWithoutA[j].word)
	})
	//combine both slices
	newWords = append(newWords, wordsWithoutA...)
	//remove the counters
	for i := 0; i < len(newWords); i++ {
		item := newWords[i]
		finalWord := item.word
		finalWords = append(finalWords, finalWord)
	}
	fmt.Println(finalWords)
}
