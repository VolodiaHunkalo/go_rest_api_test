package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

type WriteHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *WriteHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *WriteHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	if e == nil {
		panic("Logger not initialized")
	}
	return &Logger{e}
}

func (l *Logger) GetLoggerWithField(key string, value interface{}) *Logger {
	return &Logger{l.WithField(key, value)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
		DisableColors: true,
		FullTimestamp: true,
	})

	err := os.MkdirAll("logs", 0755)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("logs/all.logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&WriteHook{
		Writer:    []io.Writer{file, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	wrt := io.MultiWriter(os.Stdout, file)
	log.SetOutput(wrt)
	log.Println("Orders API Called")

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
