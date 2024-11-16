package logfile

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type LogFile struct {
	f           *os.File
	fRwm        sync.RWMutex
	size        int
	ctx         *ctx
	pathAndName string
	path        string
}

func (l *LogFile) getPrefix() string {
	return l.ctx.namePrefix
}

func NewLogFile(filePath string, opts ...Opt) (*LogFile, error) {
	c := newDefaultCtx()
	for _, opt := range opts {
		opt(c)
	}

	lf := &LogFile{
		ctx:         c,
		pathAndName: filePath,
	}

	name := ""
	pIdx := strings.LastIndexByte(filePath, '/')
	if pIdx == -1 {
		lf.path = "./"
		name = filePath
	} else {
		if pIdx == len(filePath) {
			panic(fmt.Sprintf("NewLogFile need a file path, but got a direction path: %s,"+
				"maybe you can try with [%s]", filePath, filePath[:len(filePath)-1]))
		}
		lf.path = filePath[:pIdx]
		name = filePath[pIdx+1:]
	}

	// 检查 lf.path 目录是否存在 不存在就创建
	if _, err := os.Stat(lf.path); os.IsNotExist(err) {
		err = os.MkdirAll(lf.path, 0755)
		if err != nil {
			return nil, err
		}
	}

	dotIdx := strings.LastIndexByte(name, '.')
	if dotIdx == 0 {
		panic(". can not as log fail name")
	}
	if lf.ctx.namePrefix == "" {
		if dotIdx == -1 {
			lf.ctx.namePrefix = name
		} else {
			lf.ctx.namePrefix = name[:dotIdx]
		}
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	lf.f = f

	return lf, nil
}

func (l *LogFile) trySplit() {
	if !l.ctx.isSplit() {
		return
	}
	if l.size < l.ctx.splitSize {
		return
	}

	// close old file
	l.fRwm.Lock()
	defer l.fRwm.Unlock()
	err := l.f.Close()
	if err != nil {
		fmt.Printf("LogFile: old log file can not closed: name: %s, err: %s\n",
			l.f.Name(), err.Error())
		return
	}

	// rename
	now := time.Now().Format(time.RFC3339Nano)
	newName := l.ctx.namePrefix + "_" + now + ".log"
	err = os.Rename(l.pathAndName, path.Join(l.path, newName))
	if err != nil {
		fmt.Printf("LogFile: rename fail: %s\n", err.Error())
		return
	}

	// open new file
	f, err := os.OpenFile(l.pathAndName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("LogFile: new file open fail: %s\n", err.Error())
		return
	}

	// reset cur file
	l.f = f
	l.size = 0
}

func (l *LogFile) Write(b []byte) (int, error) {
	l.trySplit()

	l.fRwm.RLock()
	n, err := l.f.Write(b)
	l.size += n
	l.fRwm.RUnlock()

	l.trySplit()
	return n, err
}
