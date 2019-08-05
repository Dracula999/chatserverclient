package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

type Client struct {
	name   string
	clConn net.Conn
}

var clients []Client

func main() {
	listener, err := net.Listen("tcp", ":3333")
	handleError(err)
	defer listener.Close()
	clientNum := 1
	for {
		conn, err := listener.Accept()
		cl := Client{name: strconv.Itoa(clientNum), clConn: conn}
		greetMsg := "Client " + cl.name + " has joined chat.\n"
		for _, client := range clients {
			client.clConn.Write([]byte(greetMsg))
		}
		clients = append(clients, cl)
		fmt.Println(len(clients))
		handleError(err)
		go manageClient(conn)
	}
}

func manageClient(conn net.Conn) {
	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Someone left")
			break
		}
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		for _, client := range clients {
			client.clConn.Write([]byte(message + "\n"))
		}
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
}
