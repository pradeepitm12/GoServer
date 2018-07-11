package utils

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"github.com/pradeepitm12/GoServer/redisUtil"

)

func HandleConnection(conn net.Conn) {
redisInstance := redisUtil.NewRedis()
redisInstance.Hset("Pradeep","h1",[]byte("Value"))
redisInstance.Hdel("Pradeep",[]string{"h1"})
	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}()

	timeoutDuration := 5 * time.Millisecond
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
