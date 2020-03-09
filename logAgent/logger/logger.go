package logger

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init(logfile, loglevel string, maxAge time.Duration) (err error) {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	Log.Out = os.Stdout
	// 判断目录是否存在，不存在则创建
	dir, _ := path.Split(logfile)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			Log.Warnf("create directory failed, err:%v\n", err)
		}
	}

	// 设置日志输出格式
	Log.SetFormatter(&logrus.JSONFormatter{})
	var level logrus.Level
	switch loglevel {
	case "trace":
		level = logrus.TraceLevel
	case "debug":
		level = logrus.DebugLevel
	case "info":
		level = logrus.InfoLevel
	case "warning":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	case "fatal":
		level = logrus.FatalLevel
	case "panic":
		level = logrus.PanicLevel
	default:
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)
	//Log.Out = os.Stdout

	writer, err := rotatelogs.New(
		// 分割后的文件名称
		logfile+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logfile),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(maxAge),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	Log.AddHook(lfHook)

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	Log.Out = file
	return
}
