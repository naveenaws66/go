package main

// stringer and iota

import "fmt"

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string{
	return stateName[ss]
}

func main() {

	fmt.Println(StateIdle)
	ns := transition(StateIdle)
	fmt.Println(ns)
	fmt.Println(StateConnected)
	ns2 := transition(ns)
	fmt.Println(ns2)
	fmt.Println(testDail(StateConnected))
}

func testDail(s ServerState) ServerState {
	return s
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Suppose we check some predicates here to
		// determine the next state...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

