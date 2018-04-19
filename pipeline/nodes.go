package pipeline

import (
	"fmt"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	ch := make(chan int)
	go func() {

		for _, v := range a {
			fmt.Printf("take data %d:\n", v)
			//第一版在这 就停了，ch 关闭了
			ch <- v
		}
		//close第二版放在这就成功了，原因：放下边关得太快了

		//clsoe 代表没数据了，不要等待发了，就不会出现 deadlock
		close(ch)

	}()

	//	close(ch) 第一版加到这，上面就close了
	return ch
}

func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		sort.Ints(a)
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}
