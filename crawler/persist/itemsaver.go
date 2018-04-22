package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	//这只是打开一个通道接受
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d:%v", itemCount, item)
			itemCount++

		}
	}()
	//外头使用的时候直接打入这个out，for 过程就能接受到
	return out

}
