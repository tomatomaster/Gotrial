// Copyright © 2017 Ryutarou Ono.

/**
参考:
http://srgia.com/docs/rfc959j.html
https://github.com/YoshikiShibata/gpl/blob/master/src/ch08/ex02/ftpServer.go

*/
package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		c := new(client)
		c.conn = conn
		go handleConn(c)
	}
}

type client struct {
	conn    net.Conn
	command string
}

func handleConn(client *client) {
	//Clientに接続完了を通知.すぐに接続できない場合は120を返すが、今回はこのような場合は存在しない想定
	io.WriteString(client.conn, "220\n")

	input := bufio.NewScanner(client.conn)
	for input.Scan() {
		client.command = input.Text()
		dispatchComand(client)
	}
}

/**
5.4. コマンド・リプライのシーケンス
受け付けるコマンドと、それに対するレスポンスが解説されている。
最悪実装を行わなくても、コマンドに対して適切なレスポンスさえ返していればクライアントとの動作確認は可能なはず

6. 状態図
このシーケンスにしたがってftpコマンドが実行されるっぽい

*/
func dispatchComand(client *client) {
	switch client.command {
	case "USER":
		io.WriteString(client.conn, "230\n")
	case "PASS":
		io.WriteString(client.conn, "230\n")
	case "ACCT":
		io.WriteString(client.conn, "230\n")
	}
}
