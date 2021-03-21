package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

type State int

const (
	primary   = 0
	secondary = 1
)

func main() {
	go sender("Numbers!")
	go receiver()
	for {

	}
	var state State
	state = secondary
	var global_counter = 0

	fmt.Println("Joining as secondary")
	// TODO: set up a udp listener

	for state == secondary {
		// TODO: observe the network
		// TODO: save global value to variable
		// TODO: break if primary is not detected
	}
	// launch a new instance
	exec.Command("go run ./main.go").Run()

	// TODO: count to the network
	for {
		global_counter++
		time.Sleep(5 * time.Second)
		// TODO: broadcast(global_counter)
	}
}

func sender(message string) {
	for {

		pc, err := net.ListenPacket("udp4", "") // automatic port number
		if err != nil {
			panic(err)
		}
		defer pc.Close()

		addr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:8829")
		if err != nil {
			panic(err)
		}

		n, err := pc.WriteTo([]byte(message), addr)
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
		time.Sleep(3 * time.Second)
	}
}

func receiver() {
	pc, err := net.ListenPacket("udp4", ":8829")
	if err != nil {
		panic(err)
	}
	defer pc.Close()
	buf := make([]byte, 1024)

	for {

		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s sent this: %s\n", addr, buf[:n])
	}
}
