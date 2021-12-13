package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// エラー処理
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Start TCP Server")

	// cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// tcpの接続アドレスを作成する
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	logFatal(err)

	// リスナーを作成する
	listener, err := net.ListenTCP("tcp", tcpAddr)
	logFatal(err)
	defer listener.Close()

	fmt.Println("Start TCP Server...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping TCP Server...")
			return
		default:
			log.Println("Running TCP Server...")
			listener.SetDeadline(time.Now().Add(time.Second * 10))

			// クライアントからのコネクション情報を受け取る
			conn, err := listener.AcceptTCP()
			if err != nil {
				switch err := err.(type) {
				case net.Error:
					if err.Timeout() {
						log.Println("Tcp Listener Close")
						return
					}
					if err.Temporary() {
						log.Printf("Temporay Error: %s\n", err.Error())
						return
					}
				default:
					log.Println("Another Error!!!")
					return
				}
				logFatal(err)
			}
			// ハンドラーに接続情報を渡す
			go handler(conn)
		}

	}
}

var responseID = 0

func handler(conn *net.TCPConn) {
	defer conn.Close()
	for {
		// リクエストを受け付けたらサーバー側に「response from server」を返す
		_, err := io.WriteString(conn, "response from server\n")
		if err != nil {
			return
		}
		responseID += 1
		_, err = io.WriteString(
			conn,
			fmt.Sprintf("responseID: %d\n", responseID),
		)
		if err != nil {
			return
		}
		fmt.Println("response returned")
		time.Sleep(time.Second)
	}
}
