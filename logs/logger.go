package logs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// const logFilesQueueLength = 100

// var logFiles = map[string]*os.File{}
// var logFilesQueue = make([]string, logFilesQueueLength)
// var mu = &sync.Mutex{}

type (
	jsonLogger struct {
		logName     string
		logFilePath string
		logs        map[string]interface{}
		mu          sync.Mutex
	}

	Logger interface {
		Log(name string, data interface{})
		Flush() error
	}
)

func NewLogger(serviceType, serviceName, functionName string) Logger {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	nowStr := now.Format("20060102")

	dirName := fmt.Sprintf("log/%s/%s/%s", nowStr, serviceType, serviceName)
	err := MkDir(dirName)
	if err != nil {
		panic(err)
	}

	return &jsonLogger{
		logFilePath: fmt.Sprintf("%s/%s.json", dirName, functionName),
		logName:     fmt.Sprintf("%s%s%s", serviceName, functionName, nowStr),
		logs:        map[string]interface{}{},
		mu:          sync.Mutex{},
	}
}

func (l *jsonLogger) Log(caption string, data interface{}) {
	l.logs[caption] = data
}

func (l *jsonLogger) Flush() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	logFile, err := os.OpenFile(l.logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer logFile.Close()

	enc := json.NewEncoder(logFile)
	for k, v := range l.logs {
		k = strings.ToUpper(fmt.Sprintf("\n ======= %s \n", k))
		logFile.Write([]byte(k))
		err = enc.Encode(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// func (l *jsonLogger) Flush() error {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	logFileLength := len(logFiles)
// 	logFile, exist := logFiles[l.logName]
// 	if !exist {
// 		f, err := os.OpenFile(l.logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
// 		if err != nil {
// 			return err
// 		}
// 		logFile = f
// 		logFiles[l.logName] = logFile
// 	}

// 	enc := json.NewEncoder(logFile)
// 	enc.SetIndent("", "    ")
// 	err := enc.Encode(l.logs)
// 	if err != nil {
// 		return err
// 	}
// 	if logFileLength < logFilesQueueLength {
// 		logFilesQueue[logFileLength] = l.logName
// 		return nil
// 	}
// 	rearLogName := logFilesQueue[len(logFilesQueue)-1]
// 	rearLogFile := logFiles[rearLogName]
// 	if logFile == rearLogFile {
// 		return nil
// 	}

// 	for i := 1; i < logFileLength; i++ {
// 		logFilesQueue[0], logFilesQueue[i] = logFilesQueue[i], logFilesQueue[0]
// 	}

// 	delete(logFiles, rearLogName)
// 	logFilesQueue[0] = l.logName

// 	return nil
// }
