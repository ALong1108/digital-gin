//go:build !windows
// +build !windows

package logger

import (
	"github.com/gin-gonic/gin"
	"syscall"
)

func init() {
	_ = syscall.Umask(000)
}

func fileCheck() {
	defer gin.Recovery()
	//防止多进程的并发操作
	if logInstance != nil && logInstance.logfile != nil {
		err := syscall.Flock(int(logInstance.logfile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
		if err != nil {
			return
		}
		defer syscall.Flock(int(logInstance.logfile.Fd()), syscall.LOCK_UN)
	}
	if logInstance != nil && logInstance.isSameFile() {
		logInstance.mu.Lock()
		defer logInstance.mu.Unlock()
		logInstance.updateFileName()
	}
}
