package main

import (
	"bufio"
	"fmt"
	"github.com/xiangbaoyan/study_go_test/pipeline"
	"os"
	"strconv"
)

func main() {
	//sink.in 总大小是512字节
	//p :=  createPipeline("sink.in",400,4)
	p := createNetWorkPipeline("sink.in", 400, 4)
	//time.Sleep(time.Hour)
	writeToFile(p, "sink.out")
	printFile("sink.out")
}
func printFile(fileName string) {
	file, e := os.Open(fileName)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	ch := pipeline.ReadSource(file, -1)
	count := 0
	for v := range ch {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}

}
func writeToFile(p <-chan int, fileName string) {
	file, e := os.Create(fileName)

	if e != nil {
		panic(e)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)

}

//实际上read file 也需要close 的
func createPipeline(fileName string, fileSize, chunkCount int) <-chan int {
	pipeline.Init()
	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, e := os.Open(fileName)
		if e != nil {
			panic(e)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReadSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemorySort(source))

	}
	return pipeline.MergeN(sortResults...)

}

func createNetWorkPipeline(fileName string, fileSize, chunkCount int) <-chan int {
	pipeline.Init()
	chunkSize := fileSize / chunkCount
	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, e := os.Open(fileName)
		if e != nil {
			panic(e)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReadSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)

		//这里只是简单把 数据 放在某个地址上
		pipeline.NetWorkSink(addr, pipeline.InMemorySort(source))
		sortAddr = append(sortAddr, addr)
		//sortResults = append(sortResults, pipeline.InMemorySort(source))
	}
	//return nil
	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		fmt.Println("get in add into sortResults")
		sortResults = append(sortResults, pipeline.NetWorkSource(addr))
	}

	return pipeline.MergeN(sortResults...)

}
