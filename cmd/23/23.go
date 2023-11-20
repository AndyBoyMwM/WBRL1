package main

import "fmt"

// Удалить i-ый элемент из слайса.
func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	i := 4

	fmt.Println("index del...")
	indexDel(sl, i)
	fmt.Println("append del...")
	appendDel(sl, i)
	fmt.Println("copy del...")
	copyDel(sl, i)
}

func indexDel(sl []int, i int) {
	fmt.Printf("in : удаляем %v элемент из: %v\n", i, sl)
	sl[i] = sl[len(sl)-1]
	sl = sl[:len(sl)-1]
	fmt.Printf("out: удалили %v элемент из: %v\n", i, sl)
}

func appendDel(sl []int, i int) {

	fmt.Printf("in : удаляем %v элемент из: %v\n", i, sl)
	fmt.Printf("len sl: %v\n", len(sl))
	sl = append(sl[:i], sl[i+1:]...)
	fmt.Printf("len sl: %v\n", len(sl))
	fmt.Printf("out: удалили %v элемент из: %v\n", i, sl)
}

func copyDel(sl []int, i int) {
	fmt.Printf("in : удаляем %v элемент из: %v\n", i, sl)
	copy(sl[i:], sl[i+1:])
	sl = sl[:len(sl)-1]
	fmt.Printf("out: удалили %v элемент из: %v\n", i, sl)
}
