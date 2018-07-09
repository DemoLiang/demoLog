package demoLog

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var (
	logger *log.Logger
)

func init() {
	pathes := strings.Split(os.Args[0], "/")
	appname := pathes[len(pathes)-1]
	//parts := strings.Split(exename, ".")
	logFilename := path.Dir(os.Args[0]) + "/" + appname + ".log"
	logfile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法打开日志文件:"+logFilename, err)
	}
	logger = log.New(logfile, "", 0)
}

func Log(level string, format string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		file = "unknown"
		line = 0
	}
	cols := strings.Split(file, "/")
	file = cols[len(cols)-1]
	userCnt := fmt.Sprintf(format, v...)
	demoLogS := DemoLogS{
		Level: level,
		Time:  time.Now(),
		File:  file,
		Line:  line,
		Cnt:   userCnt,
	}
	data, _ := json.Marshal(&demoLogS)
	logger.Println(string(data))
	log.Printf("%s\n", userCnt)
}

func Info(format string, v ...interface{}) {
	Log("info", format, v...)
}


func Debug(format string, v ...interface{}) {
	Log("debug", format, v...)
}

func Warn(format string, v ...interface{}) {
	Log("warn", format, v...)
}

func Error(format string, v ...interface{}) {
	Log("error", format, v...)
}

func Panic(format string, v ...interface{}) {
	Log("panic", format, v...)
}
