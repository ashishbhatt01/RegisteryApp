package logging

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

//Initialize() => initialize the logger for app
func Initialize() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal("cannot open/create log file", err.Error())
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(file)
}

//LogInfo() => log the info in file
func LogInfo(message string) {
	InfoLogger.Println(message)
}

//LogError() => log the error in file
func LogError(message string) {
	ErrorLogger.Println(message)
}

//LogWarning() => log the warning in file
func LogWarning(message string) {
	WarningLogger.Println(message)
}
