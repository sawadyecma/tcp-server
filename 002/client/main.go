package main

import (
	"bufio"
	"errors"
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
		err := sendMessage(connection)
		if err != nil {
			fmt.Printf("conn closed err: %s\n", err.Error())
			break
		}
	}
}

func sendMessage(connection net.Conn) error {
	fmt.Print("TcpClient> ")

	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() == false {
		fmt.Println("Ciao ciao!")
		return nil
	}

	text := stdin.Text()

	req := []byte(text)

	if len(req) == 0 {
		fmt.Println("[Error]message empty!")
		return nil
	}

	_, err := connection.Write(req)

	if err != nil {
		return err
	}

	var response = make([]byte, 4*1024)
	_, err = connection.Read(response)

	if err != nil {
		return err
	}

	fmt.Printf("Server> %s \n", response)

	if text == "q" {
		return errors.New("conn closed")
	}

	return nil
}
