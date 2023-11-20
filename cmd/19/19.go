package main

import (
	"fmt"
	"unicode/utf8"
)

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.
func main() {
	str := "欢迎参加实习"
	fmt.Printf("%s\n%s\n", str, reverseString(str))
}

func reverseString(s string) string {
	lenS := utf8.RuneCountInString(s)
	reverseString := make([]rune, lenS)
	i := 1
	for _, c := range s {
		reverseString[lenS-i] = c
		i++
	}
	return string(reverseString)
}
