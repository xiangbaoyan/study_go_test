package main

import (
	"bufio"
	"fmt"
	"github.com/xiangbaoyan/study_go_test/pipeline"
	"os"
)

func main() {
	const fileName = "large.in"
	const n = 100000000

	//创建文件
	//Create 默认创建的是二进制文件
	file, e := os.Create(fileName)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	o := pipeline.RandomSource(n)
	//bufio 有一个默认的buffer size
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, o)

	//用了bufio 就要flush 一下他
	writer.Flush()

	//读文件
	file, e = os.Open(fileName)
	if e != nil {
		panic(e)
	}

	res := pipeline.ReadSource(bufio.NewReader(file), -1)

	count := 0
	for v := range res {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}

}

func test1() {
	p1 := pipeline.InMemorySort(pipeline.ArraySource(5, 3, 3, 2, 1, 9, 22, 343))
	p2 := pipeline.InMemorySort(pipeline.ArraySource(9, 1, 10, 4, 5, 33))
	out := pipeline.Merge(p1, p2)
	for v := range out {
		fmt.Println(v)
	}
}
