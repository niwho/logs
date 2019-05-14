package logs

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// 默认debu输出
func init() {
	log.SetLevel(log.Level(5))
}

func InitLog(fileName string, level L, reservedDays int) {
	InitLogAdapter(fileName, reservedDays)
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.Level(level))
	log.AddHook(LogAdapterInstance)
}

type F log.Fields
type L log.Level

func commonFileds(fields F) *log.Entry {

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	if fields == nil {
		fields = map[string]interface{}{}
		return log.WithFields(log.Fields(fields))
	}
	fields["current_time"] = now.Format("2006-01-02 15:04:05.000")
	fields["pos"] = fmt.Sprintf("%s:%d", file, line)
	return log.WithFields(log.Fields(fields))
}

func Log(kvs map[string]interface{}) *log.Entry {
	return commonFileds(kvs)
}

func WithField(key string, val interface{}) *log.Entry {
	return commonFileds(F{key: val})
}

func AddHook(hook log.Hook) {
	log.AddHook(hook)
}
