package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.
func main() {
	dot1 := NewDot(5, 10)
	dot2 := NewDot(7, 15)
	fmt.Printf("расстояния между %v и %v =: %.4v\n", dot1, dot2, takeDist(dot1, dot2))
}

type SomeDots struct {
	a float64
	b float64
}

func (p *SomeDots) Get() (a, b float64) {
	return p.a, p.b
}

func NewDot(a, b float64) *SomeDots {
	newDot := &SomeDots{
		a: a,
		b: b,
	}
	return newDot
}

func takeDist(p1, p2 *SomeDots) float64 {
	a1, b1 := p1.Get()
	a2, b2 := p2.Get()
	return math.Sqrt(math.Pow(a1-a2, 2) + math.Pow(b1-b2, 2))
}
