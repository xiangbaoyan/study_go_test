package main

import (
	"fmt"
	"time"
)

func put(ch *chan int, arr []int) {
	go func() {
		for _, v := range arr {
			*ch <- v
		}
		close(*ch)
	}()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr1 := []int{
		8, 6, 10, 6, 7, 43, 11, 22, 11, 99, 78,
	}

	arr2 := []int{
		3, 2, 4, 6, 7, 43, 11, 22,
	}
	put(&ch1, arr1)
	put(&ch2, arr2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		fmt.Printf("ch1值:%d,ok1结果%v===="+"ch2值:%d,ok2结果%v\n", v1, ok1, v2, ok2)
		time.Sleep(time.Second * 2)
	}

}
