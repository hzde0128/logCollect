package main

import (
	"fmt"
	"github.com/hzde0128/logCollect/logAgent/conf"
	"github.com/hzde0128/logCollect/logAgent/etcd"
	"github.com/hzde0128/logCollect/logAgent/kafka"
	"github.com/hzde0128/logCollect/logAgent/logger"
	"github.com/hzde0128/logCollect/logAgent/taillog"
	"github.com/hzde0128/logCollect/logAgent/utils"
	"path"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
	wg  sync.WaitGroup
)

func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "conf/config.ini")
	if err != nil {
		panic(err)
	}
	// 加载日志文件
	// fmt.Println(path.Join(cfg.LogConf.FilePath, cfg.LogConf.FileName))
	err = logger.Init(path.Join(cfg.LogConf.FilePath, cfg.LogConf.FileName), cfg.LogConf.LogLevel, time.Duration(cfg.LogConf.MaxAge)*time.Hour*24)
	if err != nil {
		logger.Log.Warnf("初始化日志文件失败, err:%v\n", err)
	}
	// logger.Log.Info("初始化日志文件成功")
	// 1.初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		logger.Log.Errorf("init kafka failed ,err:%v\n", err)
		return
	}
	logger.Log.Info("init kafka success")

	// 2. 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		logger.Log.Errorf("init etcd failed, err:%v\n", err)
		return
	}
	logger.Log.Info("init etcd success.")

	// 为了实现每个logagent都拉取自己独有的配置，所以要以自己的IP地址作为区分
	ipStr, err := utils.GetOutboundIP(cfg.CenterConf.Address)
	if err != nil {
		logger.Log.Errorf("get local addr failed, err:%v\n", err)
		return
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		logger.Log.Errorf("get conf from etcd failed,err:%v\n", err)
		return
	}
	logger.Log.Infof("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		logger.Log.Debugf("index:%v value:%v\n", index, value)
	}

	// 3. 收集日志发往Kafka
	taillog.Init(logEntryConf)
	// 因为NewConfChan访问了tskMgr的newConfChan, 这个channel是在taillog.Init(logEntryConf) 执行的初始化
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外暴露的通道
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) // 哨兵发现最新的配置信息会通知上面的那个通道
	wg.Wait()
}
