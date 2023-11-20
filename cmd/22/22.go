package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.
func main() {
	a, _ := big.NewInt(0).SetString("34534234545657656786784565345345345345345", 10)
	b, _ := big.NewInt(0).SetString("2345234523345669567745623424254457546754", 10)

	fmt.Printf("bigNum1: %v\nbigNum2: %v\n\n", a, b)

	bgInt := big.NewInt(0)
	fmt.Println("mult:", bgInt.Mul(a, b))
	fmt.Println("div :", bgInt.Div(a, b))
	fmt.Println("sum :", bgInt.Add(a, b))
	fmt.Println("sub :", bgInt.Sub(a, b))
}
