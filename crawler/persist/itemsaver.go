package persist

import (
	"context"
	"github.com/pkg/errors"
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	//这只是打开一个通道接受

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d:%v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Sa ver:error saving item : %v %v", item, err)
			}
		}
	}()
	//外头使用的时候直接打入这个out，for 过程就能接受到
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("must supply item Type")
	}
	//第一个Index 就相当于存数据，相对于get
	indexService := client.Index().Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.
		Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
