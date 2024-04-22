package main

import (
	"fmt"
)

func main() {
	sliceOfInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range sliceOfInts {
		if sliceOfInts[i]%2 == 0 {
			fmt.Println(sliceOfInts[i], " is even number")
		} else {
			fmt.Println(sliceOfInts[i], " is odd number")
		}
	}
}
