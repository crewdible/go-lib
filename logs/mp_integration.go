package logs

import (
	"fmt"
	"log"
	"os"
	"time"
)

func MpWriteApiLog(clsName, fnName, content, ip string, singleFile bool) {
	now := time.Now().UTC()
	nowFmt := now.Format("20060102")
	dirName := fmt.Sprintf("logs/%s/api/%s", nowFmt, clsName)
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		// path/to/whatever does not exist
		// use MkdirAll for nested directory
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	// open log file
	fileName := fmt.Sprintf("%s/%s.json", dirName, fnName)
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set Date to the log : user log.SetFlag()
	// 0 means delete default time format
	log.SetFlags(0)

	if !singleFile {
		prefix := now.Format("2006-01-02 15:04:05")
		if ip != "" {
			prefix = fmt.Sprintf("%s (%s)", prefix, ip)
		}
		// Set custom time prefix
		log.SetPrefix(prefix)
	}

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.Println(fmt.Sprintf("%s\n", content))
}
