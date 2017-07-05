// Copyright © 2017 Ryutarou Ono.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := LimitReader(os.Stdin, 8)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type limitReader struct {
	reader io.Reader
	next   int
	end    int64
}

func (limit *limitReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	if int64(limit.next) >= limit.end {
		return 0, io.EOF
	}

	var nBytes int
	if int64(limit.next+len(p)) > limit.end {
		nBytes, _ = limit.reader.Read(p[:limit.end])
	} else {
		nBytes, _ = limit.reader.Read(p)
	}

	limit.next += nBytes
	return nBytes, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		reader: r,
		next:   0,
		end:    n,
	}
}
