package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

func main() {
	var global_counter = 0 // default

	// listener setup
	pc_listener, err := net.ListenPacket("udp4", ":42069")
	if err != nil {
		panic(err)
	}
	defer pc_listener.Close()
	buf_listener := make([]byte, 1024)

	// go watchdog to kill listener
	listen(pc_listener, buf_listener)

	// launch a new instance
	exec.Command("go run ./main.go").Run()

	// broadcast UDP setup
	pc_broadcast, err := net.ListenPacket("udp4", "") // automatic port number
	if err != nil {
		panic(err)
	}
	defer pc_broadcast.Close()

	addr_broadcast, err := net.ResolveUDPAddr("udp4", "255.255.255.255:42069")
	if err != nil {
		panic(err)
	}

	// TODO: count to the network
	message := make(chan string)
	go sender(pc_broadcast, addr_broadcast, message)
	for {
		global_counter++
		time.Sleep(5 * time.Second)
		message <- fmt.Sprint(global_counter)

	}
}

func sender(pc net.PacketConn, addr *net.UDPAddr, message <-chan string) {
	for {
		select {

		case msg := <-message:
			n, err := pc.WriteTo([]byte(msg), addr)
			if err != nil {
				panic(err)
			}
			_ = n // supress var not used warning
		}
	}
}

func listen(pc net.PacketConn, buf []byte) {
	// TODO: save the read value to var in main scope

	for {

		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s sent this: %s\n", addr, buf[:n])
	}
}
