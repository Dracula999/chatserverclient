package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	name string
	conn net.Conn
}

func main() {
	listener, err := net.Listen("tcp", ":3333")
	handleError(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		handleError(err)
		go manageClient(conn)
	}
}

func manageClient(conn net.Conn) {
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := "Hello"
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
}
