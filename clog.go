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

// Default verbosity, unless Level is called
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

// Level - Receives and sets a logging level
// Valid log levels are [0-5]
func Level(logLevel int) int {
	// Sets the verbosity level to the given init value
	// and returns a good status
	if logLevel < 0 || logLevel > 5 {
		// Default to 0 (disabled) if a bad value was received
		verbosity = 0
	} else {
		verbosity = logLevel
	}
	return verbosity
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
