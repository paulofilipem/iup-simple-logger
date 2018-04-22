package logger

import (
	"flag"
	"log"
	"os"
)

var (
	DEBUG     = 0
	INFO      = 1
	NOTICE    = 2
	WARNING   = 3
	ERROR     = 4
	CRITICAL  = 5
	ALERT     = 6
	EMERGENCY = 7
)

var (
	//Log is a golang log handle
	Log *log.Logger

	//Log path full or filename (If u want to write file in the runtime folder)
	logPath = "application.log"

	//Default selected log level
	LogSelectedLevel = DEBUG
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
	if LogSelectedLevel <= DEBUG {
		Log.Println("DEBUG: " + message)
	}
}

//Info func
func Info(message string) {
	if LogSelectedLevel <= INFO {
		Log.Println("INFO: " + message)
	}
}

//Notice func
func Notice(message string) {
	if LogSelectedLevel <= NOTICE {
		Log.Println("NOTICE: " + message)
	}
}

//Warning func
func Warning(message string) {
	if LogSelectedLevel <= WARNING {
		Log.Println("WARNING: " + message)
	}
}

//Error func
func Error(message string) {
	if LogSelectedLevel <= ERROR {
		Log.Println("ERROR: " + message)
	}
}

//Critical func
func Critical(message string) {
	if LogSelectedLevel <= CRITICAL {
		Log.Println("CRITICAL: " + message)
	}
}

//Alert func
func Alert(message string) {
	if LogSelectedLevel <= ALERT {
		Log.Println("ALERT: " + message)
	}
}

//Emergency func
func Emergency(message string) {
	if LogSelectedLevel <= EMERGENCY {
		Log.Println("EMERGENCY: " + message)
	}
}
