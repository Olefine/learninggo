package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	bt := make([]byte, 10)
	file, _ := os.Open("/home/egorodov/geo.csv")
	lreader := LimitReader(file, 5)
	res, _ := lreader.Read(bt)
	fmt.Println(res)
	fmt.Println(string(bt))
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(s []byte) (int, error) {
	numberOfLines := len(strings.Split(string(s), "\n"))
	*c += LineCounter(numberOfLines)

	return numberOfLines, nil
}

type WordCounter int

func (c *WordCounter) Write(s []byte) (int, error) {
	if len(s) > 0 {
		lenOfWord, _, err := bufio.ScanWords(s, true)
		*c = *c + 1
		c.Write(s[lenOfWord:])

		return int(*c), err
	}

	return 0, nil
}

type N struct {
	writer io.Writer
	writed *int64
}

func (n N) Write(b []byte) (int, error) {
	*n.writed += int64(len(b))
	return n.writer.Write(b)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	writer := N{writer: w, writed: new(int64)}
	return writer, writer.writed
}

type LReader struct {
	Limit  int64
	Reader io.Reader
}

func (r LReader) Read(p []byte) (n int, err error) {
	return r.Reader.Read(p[:r.Limit])
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := LReader{Limit: n, Reader: r}
	return lr
}
