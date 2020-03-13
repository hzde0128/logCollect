package taillog

import (
	"context"
	"github.com/hzde0128/logCollect/logAgent/kafka"
	"github.com/hzde0128/logCollect/logAgent/logger"

	"github.com/hpcloud/tail"
)

// TailTask： 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了能实现退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init() // 根据路径去打开对应的日志
	return
}

// 初始化一个TailTask实例
func (t *TailTask) init() {
	config := tail.Config{
		ReOpen: true, // 重新打开
		Follow: true, // 是否跟随
		// Whence 0表示相对于文件的原点，1表示相对于当前偏移量，2表示相对于结束。
		Location:  &tail.SeekInfo{Offset: 0, Whence: 1}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		logger.Log.Warn("tail file failed, err:", err)
	}

	go t.run() // 直接去采集日志发送到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			logger.Log.Infof("tail task:%s_%s 结束了...\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines: // 从tailObj的通道中一行一行的读取日志数据
			// 3.2 发往Kafka
			//kafka.SendToKafka(t.topic, line.Text) // 函数调用函数
			// 先把日志数据发到一个通道中
			logger.Log.Debugf("get log data from %s success, log:%v\n", t.path, line.Text)
			kafka.SendToChan(t.topic, line.Text)
			// kafka那个包中有单独的goroutine去取日志数据发到kafka
		}
	}
}
