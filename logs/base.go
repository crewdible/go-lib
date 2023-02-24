package logs

import (
	"fmt"
	"log"
	"os"
	"time"
)

// tName is Type => "api", "consumer", etc
// TIMEZONE Asia/Jakarta
func WriteLogFile(tName, clsName, fnName, content, ip string, singleFile bool) error {
	// now := time.Now().UTC()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	nowFmt := now.Format("20060102")
	// API LOG PATH "logs/%s/api/%s"
	dirName := fmt.Sprintf("log/%s/%s/%s", nowFmt, tName, clsName)
	err := MkDir(dirName)
	if err != nil {
		return err
	}

	// open log file
	fileName := fmt.Sprintf("%s/%s.json", dirName, fnName)
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
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

	// reset log settings
	log.SetFlags(log.LstdFlags)

	log.SetPrefix("")

	log.SetOutput(os.Stdout)

	return nil
}

// Store Access Token or any other json file
func WriteOtherFile(baseDir, fnName, content, ip string) error {
	dirName := fmt.Sprintf("log/%s", baseDir)
	err := MkDir(dirName)
	if err != nil {
		return err
	}

	// open log file
	fileName := fmt.Sprintf("%s/%s.json", dirName, fnName)
	logFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()

	// Set Date to the log : user log.SetFlag()
	// 0 means delete default time format
	log.SetFlags(0)

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.Println(fmt.Sprintf("%s\n", content))

	// reset log settings
	log.SetFlags(log.LstdFlags)

	log.SetPrefix("")

	log.SetOutput(os.Stdout)

	return nil
}
