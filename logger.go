package logger

import (
  "flag"
  "os"
  "log"
)

var (
  Log   *log.Logger
)


func init() {
    // set location of log file
    //var logpath = build.Default.GOPATH + "/debuggo.log"
    var logpath = "/tmp/debuggo.log"

   flag.Parse()
   var file, err1 = os.Create(logpath)

   if err1 != nil {
      panic(err1)
   }
      Log = log.New(file, "", log.LstdFlags)
      Log.Println("LogFile : " + logpath)
}

func Debug(message string){
    Log.Println("DEBUG: " + message)
}

func Info(message string){
    Log.Println("INFO: " + message)
}

func Notice(message string){
    Log.Println("NOTICE: " + message)
}

func Warning(message string){
    Log.Println("WARNING: " + message)
}

func Error(message string){
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