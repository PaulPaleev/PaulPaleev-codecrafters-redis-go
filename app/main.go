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
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handlePong(conn)
	}
}

func handlePong(conn net.Conn) {
	defer conn.Close()

	req := make([]byte, 1024)
	_, err := conn.Read(req)

	if err != nil {
		return
	}
	conn.Write([]byte("+PONG\r\n"))
}
