package main

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/pipeline"
)

func main() {
	p := pipeline.InMemorySort(pipeline.ArraySource(5, 3, 3, 2, 1, 9, 22, 343))
	for v := range p {
		fmt.Println(v)
	}

	//time.Sleep(time.Second)
}
