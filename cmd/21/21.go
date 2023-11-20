package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.
type actionInterface interface {
	Run()
}

type Animal struct {
}

func NewAnimal() *Animal {
	return &Animal{}
}

func (e *Animal) Run() {
	fmt.Println("Runing animal")
}

type Human struct {
}

func (a *Human) Run() {
	fmt.Println("Runing human")

}

type HumanAdapter struct {
	RuningHuman *Human
}

func NewHumanAdapter() actionInterface {
	return &HumanAdapter{}
}

func (r *HumanAdapter) Run() {
	fmt.Println("Адаптер переключает видеозапись с бега человека, на бег животного ")
	r.RuningHuman.Run()
}

func main() {
	an := NewAnimal()
	hum := NewHumanAdapter()

	an.Run()
	hum.Run()
}
