package main

import "fmt"

// Declare a slice and loop through it

func main() {
//fruits := []string{"apple", "mango", "papaya"}
//fmt.Println(fruits)
//for i,v := range(fruits){
//  fmt.Println(i, v)
// slice of int
//}
nums := []int{1,20,35,45}
var sum int
for _, val := range(nums) {
	sum += val
	//fmt.Println(val)
}
fmt.Println(sum)
//range over map
fruits_map := map[string]string{"a": "apple", "b": "banana",  "o": "orange", "m": "mango"}
for k,v := range(fruits_map) {
	fmt.Printf("%s --> %s\n", k, v)
}

}
