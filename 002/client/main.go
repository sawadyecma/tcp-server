package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:10000")

	if err != nil {
		panic(err)
	}

	defer connection.Close()
	for {
		sendMessage(connection)
	}
}

func sendMessage(connection net.Conn) {
	fmt.Print("TcpClient> ")

	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() == false {
		fmt.Println("Ciao ciao!")
		return
	}

	req := []byte(stdin.Text())

	if len(req) == 0 {
		fmt.Println("[Error]message empty!")
		return
	}

	_, err := connection.Write(req)

	if err != nil {
		panic(err)
	}

	var response = make([]byte, 4*1024)
	_, err = connection.Read(response)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Server> %s \n", response)
}
