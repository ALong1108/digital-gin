package logger

import (
	"digital-gin/library"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Config struct {
	TypeID       logFileType  // 1:daily 2:roll
	Level        logLevelType // 0:DEBUG 1:INFO 2:WARN 3:ERROR 4:FATAL
	FilePath     string       // log file path
	MaxFileCount int32        // max log file count
	FileSize     int64        // a log file size
	FileUnit     uint8        // a log file size unit (1 << 10*FileUnit)B
}

var initOnce = new(sync.Once)

func (c *Config) Init() {
	if c.TypeID < daily || c.TypeID > roll || c.Level < debugLogLevel || c.Level > fatalLogLevel {
		panic(c)
		return
	}

	initOnce.Do(func() {
		logLevel = c.Level

		var fileDir = filepath.Dir(c.FilePath)
		var err = os.MkdirAll(fileDir, 0777)
		if err != nil {
			panic(err)
		}
		logInstance = &logger{
			Config: Config{
				TypeID:       c.TypeID,
				Level:        c.Level,
				FilePath:     c.FilePath,
				MaxFileCount: c.MaxFileCount,
				FileSize:     c.FileSize,
				FileUnit:     c.FileUnit,
			},
			dir:      fileDir,
			filename: filepath.Base(c.FilePath),
			mu:       new(sync.RWMutex),
		}

		if c.TypeID == daily {
			t, _ := time.Parse(library.DateFormat, time.Now().Format(library.DateFormat))
			if library.IsExist(logInstance.FilePath) {
				fileInfo, err := os.Stat(logInstance.FilePath)
				if err == nil {
					t = fileInfo.ModTime()
				}
			}
			logInstance.modTime = &t
		}

		logInstance.mu.Lock()
		defer logInstance.mu.Unlock()

		if logInstance.isSameFile() {
			logInstance.updateFileName()
		} else {
			logInstance.resetFile()
		}

		// file monitor
		go func() {
			timer := time.NewTicker(1 * time.Second)
			for {
				select {
				case <-timer.C:
					fileMonitor()
				}
			}
		}()
	})
}
