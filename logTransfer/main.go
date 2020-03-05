package main

import (
	"fmt"
	"logCollect/logTransfer/conf"
	"logCollect/logTransfer/es"

	"logCollect/logTransfer/kafka"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func main() {
	// 1.加载配置文件
	err := ini.MapTo(cfg, "conf/config.ini")
	if err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	fmt.Println("load config success")

	// 2.初始化es连接
	err = es.Init(cfg.EsConf.Address)
	if err != nil {
		fmt.Printf("init es failed, err:%v\n", err)
		return
	}
	fmt.Println("init es success")

	// 3.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Printf("init kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
}
