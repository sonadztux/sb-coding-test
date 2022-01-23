package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	wordList := [...]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	fmt.Println(findAnagramWordInList(wordList[:]))
}

func findAnagramWordInList(wordList []string) [][]string {
	anagramWords := make(map[string][]string)
	var anagramWordsInSlice [][]string

	for _, word := range wordList {
		key := ""

		wordSplitted := strings.Split(word, "")
		sort.Strings(wordSplitted)
		key += strings.Join(wordSplitted, "")
		anagramWords[key] = append(anagramWords[key], word)
	}

	for _, anagramWord := range anagramWords {
		anagramWordsInSlice = append(anagramWordsInSlice, anagramWord)
	}

	return anagramWordsInSlice
}
