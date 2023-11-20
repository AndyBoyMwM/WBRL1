package main

import (
	"fmt"
)

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).
func main() {
	// Создаем объекты
	human := &Human{name: "Флинт"}
	action := &Action{Human: *human}

	// Вызываем методы
	human.Run()  // Выводит "Флинт бежит"
	action.Run() // Выводит "Флинт бежит", вызывает метод предка
}

// Human - структура родителя
type Human struct {
	name string
}

// Run - метод в структуре родителя
func (h *Human) Run() {
	fmt.Printf("%v бежит\n", h.name)
}

// Action - потомок стукруры Human
type Action struct {
	Human
}

// Run - встраиваем метод в структуру Action
func (a *Action) Run() {
	a.Human.Run()
}
