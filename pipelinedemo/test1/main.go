package main

import (
	"bufio"
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

func BufioRead() {
	r := strings.NewReader("hello,th")
	p := bufio.NewReader(r)
	//peek 是偷窥读取的意思
	//bytes, _ := p.Peek(5)
	//fmt.Println(string(bytes))
	//
	//fmt.Println(p.Buffered())
	//ReadString 读了输出来的意思，所以buffered就剩2个
	str, _ := p.ReadString(',')
	fmt.Println(str, p.Buffered())

}
func main() {
	BufioRead()
}
