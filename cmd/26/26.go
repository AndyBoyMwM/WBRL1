package main

import (
	"fmt"
	"unicode"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {

	str := []string{"value12345", "cosmic", "very_善良", "oops", "ok", "nice"}
	for _, s := range str {
		fmt.Printf("%s - %v\n", s, isUniqSimbols(s))
	}
}

func isUniqSimbols(s string) bool {
	runeSet := make(map[rune]struct{})
	for _, rn := range s {
		rn = unicode.ToLower(rn)

		// Если такой символ уже есть во множестве, возвращаем false
		if _, ok := runeSet[rn]; ok {
			return false
		}
		// Если символ ранее не встречался, добавляем его во множество
		runeSet[rn] = struct{}{}
	}
	return true
}
