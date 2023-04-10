//go:build windows
// +build windows

package logger

import "github.com/gin-gonic/gin"

func fileMonitor() {
	defer gin.Recovery()
	if logInstance != nil && logInstance.isSameFile() {
		logInstance.mu.Lock()
		defer logInstance.mu.Unlock()
		logInstance.updateFileName()
	}
}
