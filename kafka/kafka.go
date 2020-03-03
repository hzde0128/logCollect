package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func Init(addrs []string) (err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	return
}

func SendToKafka(topic, data string) {
	// 构造⼀个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("send msg failed, err: %v\n", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
