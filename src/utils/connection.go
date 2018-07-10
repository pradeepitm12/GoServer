package utils

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func HandleConnection(conn net.Conn) {

	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}()

	timeoutDuration := 1000 * time.Millisecond
	bufReader := bufio.NewReader(conn)
	conn.Write([]byte("This is pradeep reply\n"))

	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		// Read tokens delimited by newline
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", bytes)
	}
}
