package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Hello")
	listener, err := net.Listen("tcp", "3333")
	handleError(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		handleError(err)
		go manageConn(conn)
	}
}

func manageConn(conn net.Conn) {
	msg_buf := make([]byte, 1024)
	_, err := conn.Read(msg_buf)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
}
