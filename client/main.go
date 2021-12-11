package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Start TCP Client")

	// tcp, localhost:8080でリクエストを送る
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// 10秒後接続が切れる
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	// connからレスポンスを標準出力にだす
	response(os.Stdout, conn)
	fmt.Println("End TCP Client")

}

func response(dst io.Writer, src io.Reader) {
	written, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)

		log.Fatal(err)
	}
	fmt.Printf("written: %d\n", written)
}
