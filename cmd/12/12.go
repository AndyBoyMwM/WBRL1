package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество
func main() {

	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}

	for i := range set {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
