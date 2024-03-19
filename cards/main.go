package main

// funtion with return
// loop with range
// types
// funtion with receiver

func main() {
	//cards := newDeck()
	//cards.print()
	// hand, remainCards := deal(cards, 5)
	// hand.print()
	// remainCards.print()
	// c := cards.toString()
	// fmt.Println(c)
	// cards.saveToFile("temp")
	cards := newDeckFromFile("temp")
	cards.print()
}
