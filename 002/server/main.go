package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:10000")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server running at localhost:10000")

	for {
		waitClient(listener)
	}
}

func waitClient(listener net.Listener) {
	fmt.Println("wait accept")
	connection, err := listener.Accept()

	if err != nil {
		panic(err)
	}

	fmt.Printf("conn: %#+v\n", connection)
	fmt.Println("accept!")

	go goEcho(connection)
}

func goEcho(connection net.Conn) {
	defer connection.Close()
	echo(connection)
}

func echo(connection net.Conn) {

	var buf = make([]byte, 1024)

	n, err := connection.Read(buf)
	if err != nil {
		if err == io.EOF {
			return
		} else {
			panic(err)
		}
	}

	fmt.Printf("Client> %s \n", buf)

	n, err = connection.Write(buf[:n])
	if err != nil {
		panic(err)
	}

	echo(connection)
}
