package main

import "fmt"

// basic function without return
func add(a, b int) { 
 fmt.Println(a+b)
}

//function with return statements
func multiply(a, b int) int {

	return a*b

}

func main() {

	add(5,9)
	fmt.Println(multiply(5,9))
}
