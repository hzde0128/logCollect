package taillog

import (
	"fmt"
	"github.com/hzde0128/logCollect/logAgent/etcd"
	"github.com/hzde0128/logCollect/logAgent/logger"
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

	for _, logEntry := range logEntryConf {
		//conf: *etcd.LogEntry
		//logEntry.Path： 要收集的日志文件的路径
		// 初始化的时候起了多少个tailtask 都要记下来，为了后续判断方便
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		taskMgr.taskMap[mk] = tailObj
	}
	go taskMgr.run()
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理

func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 新增的
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[mk] = tailObj
				}
			}
			// 2. 配置删除
			// 找出原来t.logEntry有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry { // 从原来的配置中依次拿出配置项
				isDelete := true
				for _, c2 := range newConf { // 去新的配置中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					//t.tskMap[mk] ==> tailObj
					t.taskMap[mk].cancelFunc()
				}
			}
			logger.Log.Debug("新的配置来了！", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数，向外暴露taskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return taskMgr.newConfChan
}
