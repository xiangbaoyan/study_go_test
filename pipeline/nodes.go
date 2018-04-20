package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func ArraySource(a ...int) <-chan int {
	ch := make(chan int, 1024)
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

//相当于重新组装chan

//性能分析 都要过merge节点 ，性能反倒降低
func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read Done:", time.Now().Sub(startTime))

		sort.Ints(a)

		fmt.Println("InMemorySort Done:", time.Now().Sub(startTime))
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}
func Merge(c1, c2 <-chan int) <-chan int {
	fmt.Println("Get into Merge:")
	out := make(chan int)
	go func() {

		v1, ok1 := <-c1
		v2, ok2 := <-c2

		for ok1 || ok2 {

			if !ok2 || (ok1 && v1 <= v2) {

				out <- v1
				//这边只是v1获取新值了，原先的v2没动参与下次比较
				v1, ok1 = <-c1

			} else {
				out <- v2
				v2, ok2 = <-c2

			}
		}
		fmt.Println("Get into Merge2:", v1)

		close(out)

		fmt.Println("Merge Done:", time.Now().Sub(startTime))
	}()
	return out
}

//演示从reader 读数据
func ReadSource(r io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesSize := 0
		for {
			n, err := r.Read(buffer)
			if n > 0 {
				bytesSize += n
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesSize >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

//写二进制文件
func WriteSink(writer io.Writer, in <-chan int) {
	for v := range in {
		//写的时候得放在里面，
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

//重点应该看怎么用
func RandomSource(count int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

//两两归并
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	//重点 合并两个chan
	fmt.Println("begin Merge:")
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))

}
