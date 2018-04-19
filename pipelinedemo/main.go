package main

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/pipeline"
)

func main() {
	p1 := pipeline.InMemorySort(pipeline.ArraySource(5, 3, 3, 2, 1, 9, 22, 343))
	p2 := pipeline.InMemorySort(pipeline.ArraySource(9, 1, 10, 4, 5, 33))

	out := pipeline.Merge(p1, p2)
	for v := range out {
		fmt.Println(v)
	}

	//time.Sleep(time.Second)
}
