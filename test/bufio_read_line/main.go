package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("_args length less than 2")
	}
	fileName := args[1]
	file, e := os.Open(fileName)
	if e != nil {
		panic("_open file fail")
	}
	defer func() { file.Close() }()
	reader := bufio.NewReader(file)
	line := 0
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		if !isPrefix {
			line++
		}
	}
	fmt.Println(line)

}
