package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotates "github.com/lestrrat-go/file-rotatelogs"
)

type AsyncRotaLog struct {
	w       *rotates.RotateLogs
	bufChan chan []byte
	async   bool
	done    chan bool
}

func NewRotaLogWriter(o *Option) *AsyncRotaLog {
	absPath, err := filepath.Abs(filepath.Join(o.LogFolder, o.LogFile))
	if err != nil {
		panic(err)
	}

	w, err := rotates.New(
		fmt.Sprintf("%s-%s", absPath, o.FileNameDateFormat),
		rotates.WithMaxAge(o.MaxAge),
		rotates.WithRotationTime(o.RotationTime),
		rotates.WithLinkName(absPath),
	)
	if err != nil {
		panic(err)
	}

	al := &AsyncRotaLog{
		w:     w,
		async: o.Async,
	}

	if al.async {
		al.bufChan = make(chan []byte, logBufChanSize)
		al.done = make(chan bool, 1)
		go al.flushLog()
	}

	return al
}

func (al *AsyncRotaLog) flushLog() {
	for ai := range al.bufChan {
		al.w.Write(ai)
	}
	al.done <- true
}

func (al *AsyncRotaLog) Write(p []byte) (n int, err error) {
	if !al.async {
		return al.w.Write(p)
	}

	// maybe blocked here
	al.bufChan <- p
	return len(p), nil
}

func (al *AsyncRotaLog) Close() error {
	if al.async {
		close(al.bufChan)
		select {
		case <-al.done:
		case <-time.After(time.Duration(logFlushTimeout) * time.Second):
			fmt.Fprintf(os.Stderr, "log flush took loger than %d s\n", logFlushTimeout)
		}
	}
	return al.w.Close()
}
