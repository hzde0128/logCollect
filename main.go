package main

import (
	"fmt"
	"logAgent/conf"
	"logAgent/kafka"
	"logAgent/taillog"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "conf/config.ini")
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		return
	}
	// 1.初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed ,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2.初始化taillog
	err = taillog.Init(cfg.TailConf.Filename)
	if err != nil {
		fmt.Printf("init taillog failed, err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")

	// 3.执行任务
	run()
}
