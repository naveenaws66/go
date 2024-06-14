package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s)
	fmt.Println(s.getArea())
	//fmt.Println(area)
}

func main() {
	tri := triangle{base: 10, height: 12}
	sq := square{side: 17}
	printArea(tri)
	printArea(sq)
}
