package logger

import (
	"log"
	"os"
)

var (
	traceLogger   *log.Logger
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func Log(level, message string) {
	logger := getLogger(level)
	if logger != nil {
		logger.Printf(message)
	}
}

func init() {
	var output *os.File
	var err error
	if os.Getenv("LOG_FIlE") == "" {
		output = os.Stdout
	} else {
		output, err = os.OpenFile(os.Getenv("LOG_FIlE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	switch os.Getenv("LOG_LEVEL") {
	case "trace":
		traceLogger = log.New(output, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
		fallthrough
	case "debug":
		debugLogger = log.New(output, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		fallthrough
	case "info":
		infoLogger = log.New(output, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		fallthrough
	case "warning":
		warningLogger = log.New(output, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		fallthrough
	case "error":
		errorLogger = log.New(output, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	default:
	}

}

func getLogger(level string) *log.Logger {
	switch level {
	case "trace":
		return traceLogger
	case "debug":
		return debugLogger
	case "info":
		return infoLogger
	case "warning":
		return warningLogger
	case "error":
		return errorLogger
	default:
		return nil
	}
}
