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

//SetFileLog set log path full or filename (If u want to write file in the runtime folder)
func SetFileLog(filename string) {
	logPath = filename
	handler() //Change default handle using new filename
}

//GetFileLog get log path or filename
func GetFileLog() string {
	return logPath
}

func init() {
	flag.Parse()
	handler() //Get Default Handle
}

//handler output stream file
func handler() {
	var file, err1 = os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)

	if err1 != nil {
		panic(err1)
	}

	Log = log.New(file, "", log.LstdFlags)

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
