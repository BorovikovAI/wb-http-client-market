package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// rwx oct    meaning
// --- ---    -------
// 001 01   = execute
// 010 02   = write
// 011 03   = write & execute
// 100 04   = read
// 101 05   = read & execute
// 110 06   = read & write
// 111 07   = read & write & execute

// So 0644 is:
// * (owning) User: read & write
// * Group: read
// * Other: read

// разработчик логруса настаивает
// на использовании хуков для расширения функционала

// хук, чтобы в любое количество райтеров
// отправлять любое количество уровней логирования
type writerHook struct {
	Writer   []io.Writer
	LogLevel []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	// отправляем строчку во все райтеры
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevel
}

// чтобы можно было изменить логгер:
var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (logger *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{logger.WithField(k, v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		// определить в каком месте мы логируем
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			// файл в котором происходит логирование
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s: %d", filename, f.Line)
		},
		DisableColors: true,
		FullTimestamp: true,
		//ForceColors:   true,
	}

	// где будут храниться логи
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	// все в один файл
	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard) // по умолчанию - ничего никуда не писать

	l.AddHook(&writerHook{
		Writer:   []io.Writer{allFile, os.Stdout},
		LogLevel: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}

// возможные варианты райтеров:
// kafka -- info, debug
// file -- error, trace
// stdout -- warning, critical
