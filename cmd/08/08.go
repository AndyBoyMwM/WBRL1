package main

import "fmt"

// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func main() {
	var (
		inptNum        int64 = -79
		posBit, valBit uint8 = 1, 0
	)

	fmt.Printf("Число было: %d (%b), установить %v бит, на %v \n", inptNum, inptNum, posBit, valBit)
	numMov := setBit(inptNum, posBit, valBit)
	fmt.Printf("Число стало: %d (%b), тип числа %T\n", numMov, numMov, numMov)

	m := int64(1) << 1
	fmt.Printf("тип m: %T ", m)
}

func setBit(n int64, pos uint8, bit uint8) int64 {
	mov := int64(1) << pos
	switch bit {
	case 0:
		if n > 0 {
			return n &^ mov // (n | mov) ^ mov - тоже работает
		} else {
			return -(-n &^ mov) // для отрицательных чисел
		}
	case 1:
		return n | mov
	default:
		return n
	}
}
