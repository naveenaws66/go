package main

import "fmt"

// funtion with return
// loop with range 	
// types
// funtion with receiver 

func main() {
	var card = deck{"one", "two"}
	fmt.Println(card)
	lastcard()
}

func lastcard() {
	fmt.Println("lastcard")
}

