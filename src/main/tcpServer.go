package main

import (
	"fmt"
	"net"
	"os"

	util "../utils"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "9000"
	CONN_TYPE = "tcp"
)

func main() {
	listner, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listner.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go util.HandleConnection(conn)
	}
}
