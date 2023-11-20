package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Реализовать пересечение двух неупорядоченных множеств.
func main() {
	rand.Seed(time.Now().UnixNano())
	lenAr1, lenAr2 := 7+rand.Intn(3), 1+rand.Intn(8)
	randArr1, randArr2 := make([]int, lenAr1), make([]int, lenAr2)

	for i := 0; i < lenAr1; i++ {
		randArr1[i] = rand.Intn(9)
	}
	for i := 0; i < lenAr2; i++ {
		randArr2[i] = rand.Intn(9)
	}

	fmt.Println("Arr1: ", randArr1)
	fmt.Println("Arr2: ", randArr2)

	fmt.Println(junctionArr(randArr1, randArr2))
}

func junctionArr(inArr1, inArr2 []int) []int {
	var outArr []int
	set1 := make(map[int]bool, len(inArr1))
	set2 := make(map[int]bool, len(inArr2))

	for _, val := range inArr1 {
		set1[val] = true
	}

	for _, val := range inArr2 {
		set2[val] = true
	}

	for key := range set1 {
		if set2[key] {
			outArr = append(outArr, key)
		}
	}

	return outArr
}
