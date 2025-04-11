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

	for {
		req := make([]byte, 1024)
		conn.Read(req)
		fmt.Println("req: ", string(req))

		_, err := conn.Write([]byte(getPingResp()))

		if err != nil {
			break
		}
	}
}

func getEchoResp() string {
	return fmt.Sprintf("")
}

func getPingResp() string {
	return "+PONG\r\n"
}
