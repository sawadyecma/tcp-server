package main

import (
	"fmt"
	"io"
	"net"
	"strings"
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

	text := string(buf)

	if strings.Contains(text, "q") {
		fmt.Println("quit!")
		n, err = connection.Write([]byte("[conn cloesed]"))
		if err != nil {
			panic(err)
		}
		err := connection.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("[conn closed!]")
		return
	}

	res := append([]byte("<"), buf[:n]...)
	res = append(res, []byte(">")...)

	n, err = connection.Write(res)
	if err != nil {
		panic(err)
	}
	fmt.Println("[echo next]")
	echo(connection)
}
