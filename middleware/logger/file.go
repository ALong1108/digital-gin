package logger

import (
	"digital-gin/library"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type logFileType = uint8

const (
	_ logFileType = iota
	daily
	roll
)

type logger struct {
	Config

	dir       string
	filename  string
	rollIndex int32
	modTime   *time.Time
	mu        *sync.RWMutex
	logfile   *os.File
	l         *log.Logger // console, file
}

var logInstance *logger

func (lf *logger) resetFile() {
	var err error
	lf.logfile, err = os.OpenFile(lf.FilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer lf.logfile.Close()
	lf.l = log.New(io.MultiWriter(os.Stdout, lf.logfile), "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

// isSameFile 判断是否还是原来的文件，防止并行下文件已经被处理过
func (lf *logger) isSameFile() bool {
	fd, err := os.Stat(lf.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			lf.resetFile()
		}
		return false
	}
	curFd, err := lf.logfile.Stat()
	if err != nil {
		return false
	}
	if !os.SameFile(fd, curFd) {
		lf.resetFile()
		return false
	}

	switch lf.TypeID {
	case daily:
		t, _ := time.Parse(library.DateFormat, time.Now().Format(library.DateFormat))
		if t.After(*lf.modTime) {
			return true
		}
	case roll:
		if lf.MaxFileCount > 1 && lf.FileSize <= curFd.Size() {
			return true
		}
	}

	return false
}

func (lf *logger) updateFileName() {
	switch lf.TypeID {
	case daily:
		fileName := lf.FilePath + "." + lf.modTime.Format(library.DateFormat)
		if !library.IsExist(fileName) {
			err := os.Rename(lf.FilePath, fileName)
			if err != nil {
				lf.l.Println("updateFileName err", err.Error())
			}
			t, _ := time.Parse(library.DateFormat, time.Now().Format(library.DateFormat))
			lf.modTime = &t
			lf.resetFile()
		}
	case roll:
		lf.rollIndex = lf.rollCurrentIndex() + 1
		if library.IsExist(lf.FilePath + "." + strconv.Itoa(int(lf.rollIndex))) {
			_ = os.Remove(lf.FilePath + "." + strconv.Itoa(int(lf.rollIndex)))
		}
		_ = os.Rename(lf.FilePath, lf.FilePath+"."+strconv.Itoa(int(lf.rollIndex)))
		lf.resetFile()
	}
}

func (lf *logger) rollCurrentIndex() int32 {
	var maxTime int64
	for i := int32(1); i <= lf.MaxFileCount; i++ {
		filePath := lf.FilePath + "." + strconv.Itoa(int(i))
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			break
		}
		if fileInfo.ModTime().UnixNano() > maxTime {
			maxTime = fileInfo.ModTime().UnixNano()
			lf.rollIndex = i
		}
	}
	return lf.rollIndex % lf.MaxFileCount
}
