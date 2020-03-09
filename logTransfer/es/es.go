package es

import (
	"context"
	"logCollect/logTransfer/logger"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

// Init 初始化连接
func Init(address string, maxChanSize, nums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return err
	}
	logger.Log.Info("connect to es success")
	ch = make(chan *LogData, maxChanSize)
	for i := 0; i < nums; i++ {
		go sendToEs()
	}
	return
}

func SendToChan(msg *LogData) {
	ch <- msg
}

// SendToEs 获取记录存入Es中
func sendToEs() {
	// 链式操作
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				// Handle error
				logger.Log.Warn(err)
				continue
			}
			logger.Log.Debugf("Indexed %s %v to index %s, type %s\n", msg.Topic, put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
