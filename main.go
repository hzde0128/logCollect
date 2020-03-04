package main

import (
	"fmt"
	"logAgent/conf"
	"logAgent/etcd"
	"logAgent/kafka"
	"logAgent/taillog"
	"logAgent/utils"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
	wg  sync.WaitGroup
)

// func run() {
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}
// }

func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "conf/config.ini")
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		return
	}
	// 1.初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed ,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")

	// 2. 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")

	// 为了实现每个logagent都拉取自己独有的配置，所以要以自己的IP地址作为区分
	ipStr, err := utils.GetOutboundIP(cfg.CenterConf.Address)
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf("%s_%s",cfg.EtcdConf.Key, ipStr)
	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("get conf from etcd failed,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	// 3. 收集日志发往Kafka
	taillog.Init(logEntryConf)
	// 因为NewConfChan访问了tskMgr的newConfChan, 这个channel是在taillog.Init(logEntryConf) 执行的初始化
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外暴露的通道
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) // 哨兵发现最新的配置信息会通知上面的那个通道
	wg.Wait()
}
