package logger

import (
  "flag"
  "log"
	"os"
)

var (
	Log     *log.Logger
  logPath = "application.log"
)

func setFileLog(filename string) {
  logPath = filename
}

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

func Debug(message string) {
    Log.Println("DEBUG: " + message)
}

func Info(message string) {
    Log.Println("INFO: " + message)
}

func Notice(message string){
    Log.Println("NOTICE: " + message)
}

func Warning(message string){
    Log.Println("WARNING: " + message)
}

func Error(message string) {
    Log.Println("ERROR: " + message)
}

func Critical(message string){
    Log.Println("CRITICAL: " + message)
}

func Alert(message string){
    Log.Println("ALERT: " + message)
}

func Emergency(message string){
    Log.Println("EMERGENCY: " + message)
}