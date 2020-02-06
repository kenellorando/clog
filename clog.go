// clog
// a super simple go logging library
// github.com/kenellorando/clog

// Log Levels: Debug (5), Info (4), Warn (3), Error (2), Fatal (1), Disabled (0)
// The default is 5 (debug) unless verbosity is set with clog.Level

// args for Debug, Info, Warn: (module string, message string)
// args for Error, Fatal: (module string, message string, err error)

// All log messages are printed to stdout

package clog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Default application name, unless Init() overwites this.
var application = "clog"

// Default write-to-disk logging location, unless Init() overwites this.
var path = "/var/log/" + application

// Default write-to-disk option, unless Init() overwites this.
var write = true

// Default print to stdout, unless Init() overwites this.
var print = true

// Default verbosity, unless Init() overwites this.
var verbosity = 5

var err error

// LogData - Data contained within a log message
// 'Err' is only utilized in Error and Fatal level messages
type LogData struct {
	Time    string
	Level   string
	Module  string
	Message string
	Err     error
}

func Init(app string, logLevel int, write bool, print bool) {

	fmt.Printf("init call %s %v %v %v", app, logLevel, write, print)
	if app != "" {
		application = app
		path = "/var/log/" + application
	}

	if logLevel < 0 || logLevel > 5 {
		// Default to 5 (debug) if a bad value was received
		verbosity = 5
	} else {
		verbosity = logLevel
	}

	if write == false {
		write = false
	} else {
		write = true

		// Create a logging directory named after the application
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 766)
		}
		fmt.Println(err)
		// Create a logfile if it does not exist
		os.OpenFile(path+"/"+application+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	}

	if print == false {
		print = false
	} else {
		print = true
	}
}

// Debug - lowest level log
func Debug(module string, message string) {
	if verbosity == 5 {
		setLogData(timeNow(), "Debug", module, message, err)
	}
}

// Info - functional information level log
func Info(module string, message string) {
	if verbosity >= 4 {
		setLogData(timeNow(), "Info", module, message, err)
	}
}

// Warn - functional warning level log
func Warn(module string, message string) {
	if verbosity >= 3 {
		setLogData(timeNow(), "Warn", module, message, err)
	}
}

// Error - functional error level log
func Error(module string, message string, err error) {
	if verbosity >= 2 {
		setLogData(timeNow(), "Error", module, message, err)
	}
}

// Fatal - application failure level log
// Calls Exit
func Fatal(module string, message string, err error) {
	if verbosity >= 1 {
		setLogData(timeNow(), "Fatal", module, message, err)
	}
	os.Exit(1)
}

// Returns the date-time in specified format
func timeNow() string {
	dt := time.Now()
	return dt.Format("2006/01/02-15:04:05")
}

func writeToFile(logMessage string) {
	file, err := os.Stat(path + "/" + application + ".log")
	fmt.Println(err)
	// Rotate once logfile > 50 MB
	fmt.Println(file.Size())
	if file.Size() > file.Size()/1024/1024*50 {
		os.Rename(path+"/"+application+".log", path+"/"+application+"-"+timeNow()+".log")
	}

	logFile, _ := os.OpenFile(path+"/"+application+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	fmt.Fprintln(logFile, logMessage)
}

// Set data passed by log level methods
func setLogData(time string, level string, module string, message string, err error) {
	logData := LogData{
		Time:    time,
		Level:   level,
		Module:  module,
		Message: message,
		Err:     err,
	}
	var logMessage string
	if logData.Err != nil {
		logMessage = fmt.Sprintf(
			"%v [%5v][%s] %s\n%v\n",
			logData.Time,
			strings.ToUpper(logData.Level),
			strings.ToUpper(logData.Module),
			logData.Message,
			logData.Err)
	} else {
		logMessage = fmt.Sprintf(
			"%v [%5v][%s] %s\n%v\n",
			logData.Time,
			strings.ToUpper(logData.Level),
			strings.ToUpper(logData.Module),
			logData.Message,
			logData.Err)
	}

	if print == true {
		fmt.Printf("%s", logMessage)
	}
	if write == true {
		writeToFile(logMessage)
	}
}
