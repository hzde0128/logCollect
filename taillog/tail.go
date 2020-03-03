package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true}
	tailObj, err = tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
