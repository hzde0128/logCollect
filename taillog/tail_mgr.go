package taillog

import (
	"fmt"
	"logAgent/etcd"
	"time"
)

var taskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	taskMap     map[string]*TailTask  // 用于保存tailtask
	newConfChan chan []*etcd.LogEntry // 获取新的配置的通道
}

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集项配置信息保存起来
		taskMap:     make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
	}

	// 遍历配置项
	for _, logEntry := range logEntryConf {
		//conf: *etcd.LogEntry
		//logEntry.Path： 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
	go taskMgr.run()
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理

func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			// 1.配置新增
			// 2.配置删除
			// 3.配置变更
			fmt.Println("新的配置来了！", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数，向外暴露taskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return taskMgr.newConfChan
}
