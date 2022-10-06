package log

import (
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type DefaultHook struct {
}

func (h *DefaultHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultHook) Fire(entry *logrus.Entry) error {
	caller := getCaller()
	entry.Caller = caller
	if caller != nil && caller.Function != "" {
		s := strings.Split(caller.Function, ".")
		if len(s) > 0 {
			funcName := s[len(s)-1]
			entry.Data[Func] = funcName
		}
	}

	return nil
}

// getPackageName copy from logrus
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func getCaller() *runtime.Frame {
	/* skip 6 stack:
	-> mt-common/logs.fileInfo
	-> mt-common/logs.(*DefaultHook).Fire
	-> sirupsen/logrus.LevelHooks.Fire
	-> sirupsen/logrus.LevelHooks.Fire
	-> sirupsen/logrus.(*Entry).fireHooks
	-> logrus.(*Entry).Log
	*/
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(knownSkipFrames, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	// currentPkg := reflect.TypeOf(DefaultHook{}).PkgPath()
	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)
		// skip base component
		if strings.Contains(pkg, "sirupsen/logrus") || strings.Contains(pkg, currentPkg) {
			continue
		}
		return &f
	}
	return nil
}
