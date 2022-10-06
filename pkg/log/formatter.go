package log

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type DefaultFormatter struct {
	TextFormatter *logrus.TextFormatter
	FmtTpl        string
	FileInfoField string
	Color         bool
}

func (df *DefaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b = &bytes.Buffer{}
	var caller string
	var fieldStr string

	levelText := strings.ToUpper(entry.Level.String())
	if !df.TextFormatter.DisableLevelTruncation {
		levelText = levelText[0:4]
	}

	for k, v := range entry.Data {
		fieldStr = fieldStr + fmt.Sprintf("%s=%v ", k, v)
	}
	_, caller = defaultCallerFormat(entry.Caller)

	var s = fmt.Sprintf("%-s [%4s] "+df.FmtTpl, entry.Time.Format(df.TextFormatter.TimestampFormat), levelText, caller, entry.Message)
	if df.Color {
		levelColor := levelColor(Level(entry.Level))
		s = fmt.Sprintf("%-s \u001B[%dm[%s]\u001B[0m "+df.FmtTpl, entry.Time.Format(df.TextFormatter.TimestampFormat), levelColor, levelText, caller, entry.Message)
	}

	if fieldStr != "" {
		s = fmt.Sprintf("%s |%s\n", strings.TrimRight(s, "\n"), strings.TrimRight(fieldStr, " "))
	}

	b.WriteString(s)

	return b.Bytes(), nil
}

func levelColor(level Level) uint32 {
	var levelColor uint32
	switch level {
	case DEBUG:
		levelColor = gray
	case WARN:
		levelColor = yellow
	case ERROR, FATAL, PANIC:
		levelColor = red
	default:
		levelColor = blue
	}
	return levelColor
}

func defaultCallerFormat(f *runtime.Frame) (string, string) {
	if f == nil {
		//return "...", "???:0"
		return "", ""
	}
	s := strings.Split(f.Function, ".")
	funcName := s[len(s)-1]
	file := f.File
	file = filepath.Base(file)
	//return funcName, fmt.Sprintf("%s:%d", filepath.Base(f.File), f.Line)
	return funcName, fmt.Sprintf("%s:%d", file, f.Line)
}
