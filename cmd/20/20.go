package main

import (
	"fmt"
	"regexp"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
func main() {
	str := "欢迎。 到    我的实习  我真的很喜欢实习"
	fmt.Printf("%s\n%s\n", str, reverseWords(str))
}

func reverseWords(s string) string {
	var rev string
	reWords, _ := regexp.Compile(`(\S)+`)
	words := reWords.FindAllString(s, -1)
	reSpaces, _ := regexp.Compile(`(\s)+`)
	spaces := reSpaces.FindAllString(s, -1)
	lenWrds := len(words)
	for i := 0; i < lenWrds; i++ {
		rev += words[lenWrds-i-1]
		if i < lenWrds-1 {
			rev += spaces[lenWrds-i-2]
		}
	}
	return rev
}
