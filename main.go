package main

import (
	"fmt"
)

type State int

const (
	primary   = 0
	secondary = 1
)

func main() {
	var state State
	state = secondary

	fmt.Println("Joining as secondary")
	// set up a udp listener
	for state == secondary {
		// observe the network
	}

}
