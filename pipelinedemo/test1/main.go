package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/**
num 指定缓存区大小
*/
func ReadFrom(r io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := r.Read(p)
	if n > 0 {
		return p[:n], err
	}
	return nil, err

}

func SimpleRead() {
	bytes, _ := ReadFrom(strings.NewReader("bbbcccaaaddd你好么"), 12)
	fmt.Println(bytes)

}

func SimpleReadFromStdIn() {
	fmt.Println("read from stdin")
	bytes, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(bytes)

}
func main() {
	SimpleReadFromStdIn()
}
