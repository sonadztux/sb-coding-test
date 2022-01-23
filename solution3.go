package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket("sonadztux"))
	fmt.Println(findFirstStringInBracket("(sonadztux"))
	fmt.Println(findFirstStringInBracket("sonadztux)"))
	fmt.Println(findFirstStringInBracket(")sonadztux("))
	fmt.Println(findFirstStringInBracket(")(sonadztux"))
	fmt.Println(findFirstStringInBracket("(sonadztux)"))
	fmt.Println(findFirstStringInBracket("sonadz(sndztx)tux"))
	fmt.Println(findFirstStringInBracket("sonadztux(sndztx)"))
}

// REFACTORED CODE
func findFirstStringInBracket(str string) string {
	indexOpenBracket := strings.Index(str, "(")

	if indexOpenBracket < 0 {
		return ""
	}

	indexOpenBracket++

	indexClosingBracket := strings.Index(str[indexOpenBracket:], ")")

	if indexClosingBracket < 0 {
		return ""
	}

	return str[indexOpenBracket : indexOpenBracket+indexClosingBracket-1]
}

// ORIGINAL CODE
// func findFirstStringInBracket(str string) string {
// 	if len(str) > 0 {
// 		indexFirstBracketFound := strings.Index(str, "(")
// 		if indexFirstBracketFound >= 0 {
// 			runes := []rune(str)
// 			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
// 			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
// 			if indexClosingBracketFound >= 0 {
// 				runes := []rune(wordsAfterFirstBracket)
// 				return string(runes[1 : indexClosingBracketFound-1])
// 			} else {
// 				return ""
// 			}
// 		} else {
// 			return ""
// 		}
// 	} else {
// 		return ""
// 	}
// 	return ""
// }
