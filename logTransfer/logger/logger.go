package logger

import (
	"os"
	"path"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	// Log 全局日志
	Log = logrus.New()
)

// logLevel 设置日志等级
func logLevel(loglevel string) (level logrus.Level) {
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
	return
}

// 创建目录
func mkDir(logfile string) (logpath string, err error) {
	logpath, err = filepath.Abs(logfile)
	if err != nil {
		return
	}
	dir, _ := path.Split(logpath)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}
	}
	return
}

// Init 日志初始化
func Init(logfile, loglevel string, maxAge time.Duration) (err error) {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	Log.Out = os.Stdout
	
	// 判断目录是否存在，不存在则创建
	logpath, err := mkDir(logfile)
	if err != nil {
		Log.Errorf("创建目录失败, err:%v\n", err)
		return
	}

	// 设置日志输出格式
	Log.SetFormatter(&logrus.JSONFormatter{})

	// 获取日志级别
	Log.SetLevel(logLevel(loglevel))

	writer, err := rotatelogs.New(
		// 分割后的文件名称
		logpath+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logpath),
		// 设置最大保存时间
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

	// 日志输出到文件
	file, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	Log.Out = file
	return
}
