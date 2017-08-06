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

実行のサンプル
> ftp
> open localhost 8080
> ls
> get [filename]
> send [filename]
> exit
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/**
This FTP server runs 8080 port.
How to access
$ftp
> open localhost 8080
*/
func main() {
	if err := os.Chdir(os.Getenv("HOME")); err != nil {
		checkError(err)
	}
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
	conn      net.Conn //Communication Connection
	dconn     net.Conn //Data Connection
	command   []string //Command ex) cd [dir]
	transMode string   //ascii or binary
	wdir      string   //current dir
}

//See http://tooljp.com/qa/622E6C4723FB1D0C49257C4A005DDCD1.html
type asciiMode struct {
	w io.Writer
}

func (a *asciiMode) Write(p []byte) (n int, err error) {
	buf := make([]byte, 0, len(p))
	for _, b := range p {
		if b == '\n' {
			buf = append(buf, '\r')
		}

		if b != '\r' {
			buf = append(buf, b)
		}
	}
	return a.w.Write(buf)
}

const (
	//Mode
	ascii  = "ascii"
	binary = "binary"

	//Status
	cOK             = "200"
	cNotImplement   = "502"
	dConnectOpened  = "125" //125 Data connection already open; transfer starting.
	dConnectClosing = "226" //226 Closing data connection.
	fActionNotTaken = "450" //450 Requested file action not taken.
)

func (c *client) writeStatus(statusCode string) {
	io.WriteString(c.conn, fmt.Sprintln(statusCode))
}

func (c *client) closeConn() {
	if c.conn != nil {
		c.conn.Close()
	}
	if c.dconn != nil {
		c.dconn.Close()
	}
}

func handleConn(client *client) {
	defer client.closeConn()
	client.writeStatus("220") //Clientに接続完了を通知.すぐに接続できない場合は120を返すが、今回はこのような場合は存在しない想定
	client.transMode = ascii
	var err error
	if client.wdir, err = os.Getwd(); err != nil {
		checkError(err)
	}
	input := bufio.NewScanner(client.conn)
	for input.Scan() {
		client.command = strings.Split(input.Text(), " ")
		dispatchComand(client)
	}
}

/**
1: クライアントからのコマンドを受け取る
2: 必要に応じてクライアントに接続状態を返す
3: コマンドに応じた処理を行う
4: 処理結果をクライアントに通知する
*/
func dispatchComand(client *client) {
	switch client.command[0] {
	case "USER":
		client.writeStatus("331")
	case "QUIT":
		client.writeStatus("230")
		client.closeConn()
	case "PORT": //For Windows
		portComm(client)
	case "TYPE":
		typeComm(client)
	case "CWD":
		cwdComm(client)
	case "PWD":
		client.writeStatus("257")
	case "MODE":
		fmt.Println("[DEBUG] Not Implement Yet")
	case "STRU":
		fmt.Println("[DEBUG] Not Implement Yet")
	case "RETR":
		retrComm(client)
	case "STOR":
		storComm(client)
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
	case "SYST": //For Windows https://tools.ietf.org/html/rfc1700 OPERATING SYSTEM NAMES 参照
		client.writeStatus("215 OSX system type")
	default:
		client.writeStatus(cNotImplement)
		fmt.Printf("[DEBUG] NoSupport Commnad: %s \n", client.command)
	}
}

/**
Commnads
*/

func storComm(client *client) {
	defer client.dconn.Close()
	client.writeStatus(dConnectOpened)
	file, err := os.Create(client.command[1])
	if err != nil {
		client.writeStatus(fActionNotTaken)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, client.dconn)
	if err != nil {
		client.writeStatus(fActionNotTaken)
		return
	}
	client.writeStatus(dConnectClosing)
}

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

func portComm(client *client) {
	addr := parsePORTAddr(client.command[1])
	fmt.Printf("[DEBUG] Connect to %s \n", addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	checkError(err)
	client.dconn, err = net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	client.writeStatus(cOK)
}

func listComm(client *client) {
	client.writeStatus(dConnectOpened)
	cmd := exec.Command("ls", "-la")
	reader, err := cmd.StdoutPipe()
	checkError(err)
	aMode := asciiMode{client.dconn}
	go io.Copy(&aMode, reader)
	cmd.Start()
	cmd.Wait()
	client.dconn.Close()
	client.writeStatus(dConnectClosing)
}

func retrComm(client *client) {
	defer client.dconn.Close()
	client.writeStatus(dConnectOpened)
	file, err := os.Open(client.command[1])
	if err != nil {
		client.writeStatus(fActionNotTaken)
		return
	}
	switch client.transMode {
	case ascii:
		asciimode := asciiMode{client.dconn}
		io.Copy(&asciimode, file)
	case binary:
		io.Copy(client.dconn, file)
	default:
	}
	client.writeStatus(dConnectClosing)
}

func cwdComm(client *client) {
	var err error
	if err = os.Chdir(client.command[1]); err != nil {
		client.writeStatus("550")
		return
	}
	if client.wdir, err = os.Getwd(); err != nil {
		client.writeStatus("550")
		return
	}
	client.writeStatus("250")
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

func parsePORTAddr(addr string) string {
	a := strings.Split(addr, ",")
	ip := strings.Join(a[0:3], ".")
	portA, _ := strconv.Atoi(a[4])
	portB, _ := strconv.Atoi(a[5])
	port := portA*256 + portB
	return fmt.Sprintf("%s:%s", ip, port)
}
