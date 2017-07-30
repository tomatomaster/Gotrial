// Copyright © 2017 Ryutarou Ono.

/**
方針
RFCは読んでもいまいちわからないので補足程度に
サンプル実装とクライアント側の実装から仕様を確認する




参考:
http://srgia.com/docs/rfc959j.html
https://github.com/YoshikiShibata/gpl/blob/master/src/ch08/ex02/ftpServer.go
http://www.geekpage.jp/programming/linux-network/tcp-1.php [TCPを使う]
http://x68000.q-e-d.net/~68user/net/ftp-1.html[FTP クライアントを作ってみよう (1)
]

動作確認用にMacでFTPサーバをローカル実行する
http://osxdaily.com/2011/09/29/start-an-ftp-or-sftp-server-in-mac-os-x-lion/
以下のaliasを登録した
alias startftp="sudo -s launchctl load -w /System/Library/LaunchDaemons/ftp.plist"
alias stopftp="sudo -s launchctl unload -w /System/Library/LaunchDaemons/ftp.plist"
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"strings"
)

/**
This FTP server runs 8080 port.
How to access ?
$ftp
> open localhost 8080
*/
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	err = checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		checkError(err)
		c := new(client)
		c.conn = conn
		go handleConn(c)
	}
}

type client struct {
	conn      net.Conn
	dconn     net.Conn
	command   []string
	transMode string
}

const (
	//Mode
	ascii  = "ascii"
	binary = "binary"

	//Status
	cOK           = "200"
	cNotImplement = "502"
)

func (c *client) writeStatus(statusCode string) {
	io.WriteString(c.conn, fmt.Sprintln(statusCode))
}

func handleConn(client *client) {
	//Clientに接続完了を通知.すぐに接続できない場合は120を返すが、今回はこのような場合は存在しない想定
	client.writeStatus("220")
	client.transMode = ascii //デフォルトの転送モード(変更しても良い)
	input := bufio.NewScanner(client.conn)
	for input.Scan() {
		client.command = strings.Split(input.Text(), " ")
		fmt.Printf("[DEBUG] Command: %s\n", client.command)
		dispatchComand(client)
	}
}

/**
5.4. コマンド・リプライのシーケンス
受け付けるコマンドと、それに対するレスポンスが解説されている。
最悪実装を行わなくても、コマンドに対して適切なレスポンスさえ返していればクライアントとの動作確認は可能なはず

6. 状態図
このシーケンスにしたがってftpコマンドが実行されるっぽい


 USER, QUIT, PORT,
                    TYPE, MODE, STRU,(デフォルト値のためのもの)
                    RETR, STOR,
                    NOOP
*/
func dispatchComand(client *client) {
	switch client.command[0] {
	case "USER":
		client.writeStatus("331")
	case "QUIT":
		client.writeStatus("230")
	case "PORT":
		fmt.Println("[DEBUG] Not Implement Yet")
	case "TYPE":
		typeComm(client)
	case "MODE":
		fmt.Println("[DEBUG] Not Implement Yet")
	case "STRU":
		fmt.Println("[DEBUG] Not Implement Yet")
	case "RETR":
		fmt.Println("[DEBUG] Not Implement Yet")
	case "STOR":
		fmt.Println("[DEBUG] Not Implement Yet")
	case "NOOP":
		fmt.Println("[DEBUG] Not Implement Yet")
	case "EPRT":
		eprtComm(client)
	case "PASS":
		client.writeStatus("530")
	case "ACCT":
		client.writeStatus("530")
	case "LIST":
		listComm(client)
	default:
		client.writeStatus(cNotImplement)
		fmt.Printf("[DEBUG] NoSupport Commnad: %s \n", client.command)
	}
}

/**
Commnads
*/

func ls(params ...string) {
	exec.Command("ls", params...)
}

func typeComm(client *client) {
	switch client.command[1] {
	case "A":
		fmt.Println("[DEBUG] ASCII Mode")
		client.transMode = ascii
	case "I":
		client.transMode = binary
		fmt.Println("[DEBUG] BINARY Mode")
	default:
		fmt.Printf("Unsupported Mode: %s", client.command[1])
	}
	client.writeStatus(cOK)
}

func eprtComm(client *client) {
	addr := parseEPRTAddr(client.command[1])
	fmt.Printf("[DEBUG] Connect to %s \n", addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	checkError(err)
	client.dconn, err = net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	client.writeStatus(cOK)
}

func listComm(client *client) {
	client.writeStatus("125")
	cmd := exec.Command("ls", "-la")
	reader, err := cmd.StdoutPipe()
	checkError(err)
	go io.Copy(client.dconn, reader)
	cmd.Start()
	cmd.Wait()
	client.writeStatus("226")
}

//Utils
func checkError(err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return err
}

/**
EPRT command sample
EPRT |1|132.235.1.2|6275|
EPRT |2|1080::8:800:200C:417A|5282|
EPRT |2|::1|53287|

|1| IPv4 |2| IPv6

http://www5d.biglobe.ne.jp/stssk/rfc/rfc2428j.html
*/
func parseEPRTAddr(addr string) string {
	a := strings.Split(addr, "|")
	switch a[1] {
	case "1":
		addr = fmt.Sprintf("%s:%s", a[2], a[3])
	case "2":
		addr = fmt.Sprintf("[%s]:%s", a[2], a[3])
	default:
		fmt.Errorf("Unsupported Type %s\n", a[1])
	}
	return addr
}
