package main

import "fmt"

type rect struct {
	length int
	breadth int
}

func (r *rect) area() int{
	return r.length * r.breadth
}

func (r rect) perim() int {
	return 2 * r.length + 2 *  r.breadth
}

func main() {
	r := rect{ length: 5, breadth: 10}
	fmt.Println(r.area())
	fmt.Println(r.perim())
	
	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())

}
