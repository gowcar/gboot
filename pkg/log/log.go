package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Level uint32

const (
	PANIC Level = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
)

const (
	maximumCallerDepth int    = 25
	knownSkipFrames    int    = 8
	logBufChanSize     int    = 10000
	logFlushTimeout    int    = 10 // second
	currentPkg         string = "gboot/pkg/log"
)

const (
	FileNameDateFormat = "%Y%m%d"
	TimestampFormat    = "2006-01-02 15:04:05.000"
	Text               = "text"
	JSON               = "json"
	DataKey            = "data"
	Func               = "func"
)

var defaultLogger *logrus.Logger
var defaultWriter io.WriteCloser

type Option struct {
	// log目录
	LogFolder string
	// log文件名
	LogFile string
	// 日志类型 json|text
	LogType string
	// 文件名的日期格式
	FileNameDateFormat string
	// 日志中日期时间格式
	TimestampFormat string
	// 日志级别
	LogLevel string
	// 是否打印到控制台
	Log2Console bool
	// 是否异步刷新
	Async bool
	// 日志最长保存多久
	MaxAge time.Duration
	// 日志默认多长时间轮转一次
	RotationTime time.Duration
	// 是否开启记录文件名和行号
	IsEnableRecordFileInfo bool
	// 文件名和行号字段名
	FileInfoField string
	// json日志是否美化输出
	JSONPrettyPrint bool
	// json日志条目中 数据字段都会作为该字段的嵌入字段
	JSONDataKey string
	// Text模式下日志打印的格式字符串，两个占位符：caller、message内容
	TextFmtTpl string
}

func DefaultOption() *Option {
	return &Option{
		LogFolder:              "/tmp/logs/",
		LogFile:                "messages.log",
		LogType:                "text",
		FileNameDateFormat:     "%Y-%m-%d",
		TimestampFormat:        "2006-01-02 15:04:05.000",
		LogLevel:               "debug",
		Log2Console:            true,
		Async:                  false,
		MaxAge:                 7 * 24 * time.Hour,
		RotationTime:           24 * time.Hour,
		IsEnableRecordFileInfo: true,
		FileInfoField:          "caller",
		JSONPrettyPrint:        false,
		JSONDataKey:            "json",
		// caller, message
		TextFmtTpl: "%-s %s\n",
	}
}
func DoInit() {
	InitLogger(DefaultOption())
}

func InitLogger(option *Option) {
	if err := makeDirAll(option.LogFolder); err != nil {
		panic(err)
	}

	defaultLogger = logrus.New()
	defaultLogger.SetOutput(ioutil.Discard)
	level, err := logrus.ParseLevel(option.LogLevel)
	if err != nil {
		panic(err)
	}
	defaultLogger.Level = level

	switch option.LogType {
	case JSON:
		format := &logrus.JSONFormatter{
			TimestampFormat: option.TimestampFormat,
			PrettyPrint:     option.JSONPrettyPrint,
		}
		if option.JSONDataKey != "" {
			format.DataKey = option.JSONDataKey
		}
		defaultLogger.Formatter = format
	default:
		defaultLogger.Formatter = &DefaultFormatter{
			TextFormatter: &logrus.TextFormatter{
				TimestampFormat: option.TimestampFormat,
			},
			Color:         true,
			FileInfoField: option.FileInfoField,
			FmtTpl:        option.TextFmtTpl,
		}
	}
	if option.Log2Console {
		defaultLogger.Out = os.Stdout
	}

	defaultWriter = NewRotaLogWriter(option)

	fileFormatter := &DefaultFormatter{
		TextFormatter: &logrus.TextFormatter{
			TimestampFormat: option.TimestampFormat,
		},
		Color:         false,
		FileInfoField: option.FileInfoField,
		FmtTpl:        option.TextFmtTpl,
	}

	fileHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: defaultWriter,
		logrus.InfoLevel:  defaultWriter,
		logrus.WarnLevel:  defaultWriter,
		logrus.ErrorLevel: defaultWriter,
		logrus.FatalLevel: defaultWriter,
		logrus.PanicLevel: defaultWriter,
	}, fileFormatter)

	defaultLogger.Hooks.Add(&DefaultHook{})
	defaultLogger.Hooks.Add(fileHook)
}

func makeDirAll(logPath string) error {
	logDir := path.Dir(logPath)
	_, err := os.Stat(logDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
				return fmt.Errorf("create <%s> error: %s", logDir, err)
			}
			return nil
		}
		return err
	}

	return nil
}

func Flush() error {
	if defaultWriter != nil {
		defaultWriter.Close()
	}
	return defaultLogger.Writer().Close()
}

func Writer() io.Writer {
	return defaultLogger.Writer()
}

func Debug(message string, args ...interface{}) {
	if args != nil {
		defaultLogger.Logf(logrus.DebugLevel, message, args...)
	} else {
		defaultLogger.Debug(message)
	}
}

func Info(message string, args ...interface{}) {
	if args != nil {
		defaultLogger.Logf(logrus.InfoLevel, message, args...)
	} else {
		defaultLogger.Info(message)
	}
}

func Warn(message string, args ...interface{}) {
	if args != nil {
		defaultLogger.Logf(logrus.WarnLevel, message, args...)
	} else {
		defaultLogger.Warn(message)
	}
}

func Error(message string, args ...interface{}) {
	if args != nil {
		defaultLogger.Logf(logrus.ErrorLevel, message, args...)
	} else {
		defaultLogger.Error(message)
	}
}

func Fatal(message string, args ...interface{}) {
	if args != nil {
		defaultLogger.Logf(logrus.FatalLevel, message, args...)
	} else {
		defaultLogger.Fatal(message)
	}
}

func Panic(message string, args ...interface{}) {
	if args != nil {
		defaultLogger.Logf(logrus.PanicLevel, message, args...)
	} else {
		defaultLogger.Panic(message)
	}
}

// GetCustomLog return log entry with context set by custom k,v.
// note: no need release
func GetCustomLog(k string, v interface{}) *logrus.Entry {
	return defaultLogger.WithField(k, v)
}

func StderrFatalf(message string, args ...interface{}) {
	defaultLogger.Logf(logrus.FatalLevel, message, args...)
}

func GetLogger() *logrus.Logger {
	logger := defaultLogger
	return logger
}

// SetLogLevel set log level, we need when server hot-reload for debug
func SetLogLevel(level string) {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return
	}
	defaultLogger.SetLevel(l)
}
