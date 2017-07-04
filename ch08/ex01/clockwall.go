// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var cityHost map[string]string

func main() {
	cityHost = make(map[string]string)
	for _, arg := range os.Args[1:] {
		cityHostRelation := strings.Split(arg, "=")
		if len(cityHostRelation) != 2 {
			log.Fatalf("Unexpected argument %v", arg)
		}
		city := cityHostRelation[0]
		host := cityHostRelation[1]
		cityHost[city] = host
	}

	for key, val := range cityHost {
		conn, err := net.Dial("tcp", val)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		fmt.Println(key)
		go mustCopy(os.Stdout, conn)
	}
	select {}
}

func mustCopy(dst io.Writer, src io.Reader) {
	for {
		buf := make([]byte, 1024)
		n, err := src.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		dst.Write(buf[:n])
	}
}
