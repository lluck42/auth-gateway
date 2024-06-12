package bean

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func LogInit() {

	Log = logrus.New()
	// 设置日志级别为xx以及以上
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间格式化
	})
	Log.SetReportCaller(true)

	// log文件
	file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Log.Panic("LogInit失败")
	}

	Log.Out = file
}
