package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println("in: ", arr)
	fmt.Println("out: ", sortArr(arr))
}

func sortArr(inArr []float64) map[int][]float64 {
	outArr := make(map[int][]float64)

	for _, v := range inArr {
		key := int(v) / 10 * 10
		outArr[key] = append(outArr[key], v)
	}
	return outArr
}
