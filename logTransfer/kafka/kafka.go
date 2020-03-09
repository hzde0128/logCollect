package kafka

import (
	"logCollect/logTransfer/es"
	"logCollect/logTransfer/logger"
	"sync"

	"github.com/Shopify/sarama"
)

// Init 初始化kafka连接，准备发送数据给es
func Init(addr []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		logger.Log.Errorf("fail to start consumer, err:%v\n", err)
		return
	}

	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		logger.Log.Errorf("fail to get list of partition:err%v\n", err)
		return
	}
	logger.Log.Debugf("topic:%s, partition:%v\n", topic, partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		var cp sarama.PartitionConsumer
		cp, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logger.Log.Errorf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		// defer cp.AsyncClose()
		// 异步从每个分区消费信息
		var wg sync.WaitGroup
		defer wg.Done()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range cp.Messages() {
				logger.Log.Debugf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 准备发送数据给es
				ld := es.LogData{Topic: topic, Data: string(msg.Value)}
				//es.SendToES(topic, ld) // 函数调用函数
				// 优化一下: 直接放到一个chan中
				es.SendToChan(&ld)
			}
		}(cp)
		wg.Wait()
	}
	return
}
