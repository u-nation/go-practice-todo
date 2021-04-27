package logger

import (
	"fmt"
	"github.com/u-nation/go-practice-todo/pkg/util"
	"runtime"
	"strconv"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var (
	logentry *logrus.Entry
)

type PlainFormatter struct {
	PackageName     string
	VerifyMode      bool
	TimestampFormat string
	LevelDesc       []string
}

func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(f.TimestampFormat)

	return []byte(fmt.Sprintf(
		"[%s] [%s] [%s] %s\n",
		timestamp,
		f.LevelDesc[entry.Level],
		f.PackageName,
		entry.Message,
	)), nil
}

func NewLogger(packageName string, isDevMode bool) error {
	formatter := &PlainFormatter{
		PackageName:     packageName,
		VerifyMode:      isDevMode,
		TimestampFormat: "2006-01-02 15:04:05.999+00:00",
		LevelDesc: []string{
			"PANC",
			"FATL",
			"ERRO",
			"WARN",
			"INFO",
			"DEBG",
		},
	}

	logrus.SetFormatter(formatter)
	logrus.SetOutput(colorable.NewColorableStdout())

	if isDevMode {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logentry = logrus.WithFields(logrus.Fields{})

	return nil
}

func Debug(requestId string, msg ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	logentry.Debug(format(filename, line, requestId, msg...))
}

func Info(requestId string, msg ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	logentry.Info(format(filename, line, requestId, msg...))
}

func Warn(requestId string, msg ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	logentry.Warn(format(filename, line, requestId, msg...))
}

func Error(requestId string, msg ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	logentry.Error(format(filename, line, requestId, msg...))
}

func format(filename string, line int, requestId string, msg ...interface{}) string {
	return fmt.Sprintf(
		"[%s:%s] [%s] [%s] %v",
		filename,
		strconv.Itoa(line),
		strconv.Itoa(util.GoroutineID()),
		requestId,
		msg,
	)
}
