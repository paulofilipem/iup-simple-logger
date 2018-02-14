package logger

import (
	"flag"
	"log"
	"os"
)

var (
	//Log is a golang log handle
	Log *log.Logger

	//Log path full or filename (If u want to write file in the runtime folder)
	logPath = "application.log"
)

//setFileLog set log path full or filename (If u want to write file in the runtime folder)
func setFileLog(filename string) {
	logPath = filename
}

//getFileLog get log path or filename
func getFileLog() string {
	return logPath
}

func init() {

	flag.Parse()
	var file, err1 = os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)

	if err1 != nil {
		panic(err1)
	}

	Log = log.New(file, "", log.LstdFlags)
	Log.Println("LogFile : " + logPath)
}

//Debug func
func Debug(message string) {
	Log.Println("DEBUG: " + message)
}

//Info func
func Info(message string) {
	Log.Println("INFO: " + message)
}

//Notice func
func Notice(message string) {
	Log.Println("NOTICE: " + message)
}

//Warning func
func Warning(message string) {
	Log.Println("WARNING: " + message)
}

//Error func
func Error(message string) {
	Log.Println("ERROR: " + message)
}

//Critical func
func Critical(message string) {
	Log.Println("CRITICAL: " + message)
}

//Alert func
func Alert(message string) {
	Log.Println("ALERT: " + message)
}

//Emergency func
func Emergency(message string) {
	Log.Println("EMERGENCY: " + message)
}
