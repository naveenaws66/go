package main

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message recevied")
	}

	msg := "hi"
	
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message received")
	}
	
	select {
	case msg1 := <-messages:
		fmt.Println("received message", msg1)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

