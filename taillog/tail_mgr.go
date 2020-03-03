package taillog

import "logAgent/etcd"

var taskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	//taskMap map[string]*TailTask
}

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntry: logEntryConf, // 把当前的日志收集项配置信息保存起来
	}

	// 遍历配置项
	for _, logEntry := range logEntryConf {
		//conf: *etcd.LogEntry
		//logEntry.Path： 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
