package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn, err := l.Accept()

	for {
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		handlePong(conn)
	}
}

func handlePong(conn net.Conn) {
	defer conn.Close()

	mess := make([]byte, 1024)
	_, err := conn.Read(mess)
	if err != nil {
		return
	}
	conn.Write([]byte("+PONG\r\n"))
}
