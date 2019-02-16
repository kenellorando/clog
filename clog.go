// clog
// a super simple go logging library
// github.com/kenellorando/clog

// Log Levels: Debug, Info, Warn, Error, Fatal
// args for Debug, Info, Warn: (module string, message string)
// args for Error, Fatal: (module string, message string, err error)

// Prints log messages to stdout

package clog

import (
	"fmt"
	"strings"
	"time"
)

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

// Debug - lowest level log
// Atomic level application logging
func Debug(module string, message string) {
	setLogData(timeNow(), "Debug", module, message, err)
}

// Info - functional information level log
// Function level application information
func Info(module string, message string) {
	setLogData(timeNow(), "Info", module, message, err)
}

// Warn - functional warning level log
// Monitoring for potential erroneous or fatal situations
func Warn(module string, message string) {
	setLogData(timeNow(), "Warn", module, message, err)
}

// Error - functional error level log
// Failure of a function or feature to execute
// The application should still be operable
func Error(module string, message string, err error) {
	setLogData(timeNow(), "Error", module, message, err)
}

// Fatal - application failure level log
// Indicates the application is inoperable, or a
// shutdown of the application is imminent
func Fatal(module string, message string, err error) {
	setLogData(timeNow(), "Fatal", module, message, err)
}

// Returns the date-time in specified format
func timeNow() string {
	dt := time.Now()
	return dt.Format("2006/01/02 15:04:05")
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

	printLogMessage(logData)
}

// Print the formatted log message to stdout
func printLogMessage(ld LogData) {
	// If there is an error set in logData,
	// Print the error with the log message
	if ld.Err != nil {
		logMessage := fmt.Sprintf(
			"%v [%5v][%s] %s\n%v\n",
			ld.Time,
			strings.ToUpper(ld.Level),
			strings.ToUpper(ld.Module),
			ld.Message,
			ld.Err,
		)
		fmt.Printf(logMessage)
	} else {
		logMessage := fmt.Sprintf(
			"%v [%-5v][%s] %s\n",
			ld.Time,
			strings.ToUpper(ld.Level),
			strings.ToUpper(ld.Module),
			ld.Message,
		)
		fmt.Printf(logMessage)
	}
}
