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

// function with multiple return values

func mul(a, b int) (int, int, int){

	return a, b, a*b

}

// variadac functions

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
		for _, v := range(nums) {
			total += v
		}
	fmt.Println(total)

}
func main() {

	add(5,9)

	fmt.Println(multiply(5,9))
	fmt.Println(mul(5,9))
	sum(1, 2, 3, 4)
}
