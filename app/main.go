package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
		var err error
		req := make([]byte, 1024)
		conn.Read(req)
		strReq := string(req)
		if strings.Contains(strReq, "ECHO") {
			arg := getEchoArg(strReq)
			_, err = conn.Write([]byte(getEchoResp(len(arg), arg)))
		} else {
			_, err = conn.Write([]byte(getPingResp()))
		}

		if err != nil {
			break
		}
	}
}

func getEchoResp(size int, arg string) string {
	return fmt.Sprintf("%v\r\n%v\r\n", size, arg)
}

func getEchoArg(strReq string) string {
	arg := strings.Split(strReq, "\r\n")[4]
	return arg
}

func getPingResp() string {
	return "+PONG\r\n"
}
