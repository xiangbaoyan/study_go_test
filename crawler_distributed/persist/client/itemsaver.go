package client

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	//这只是打开一个通道接受

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d:%v", itemCount, item)
			itemCount++

			//Call RPC TO SAVE
			result := ""
			client.Call("ItemSaverService.Save", item, &result)
			if err != nil {
				log.Printf("Item Sa ver:error saving item : %v %v", item, err)
			}
		}
	}()
	//外头使用的时候直接打入这个out，for 过程就能接受到
	return out, nil
}
