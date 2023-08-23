# ⚠️ Archived - Deprecation Notice
`clog` will no longer be maintained. I originally made this for [Cadence](https://github.com/kenellorando/cadence), to make up for Go standard log package's then-lack of logging levels and customizable information. The release of Go 1.21's standard library includes `slog`, an extremely similar (in name and function) stuctured logger which supports log levels and custom tags. See Go's [slog post](https://go.dev/blog/slog) for details. 

Still, `clog` remains viable for Go projects 1.20 and prior.

# clog
clog is a Go logging package. It is ideal for applications in environments whose standard outputs are caught by other systems, such as container runtimes and Linux system utilities. 

Log messages look like this. The log level and custom label are easy to read and parse.
```
2019/02/21 12:49:44 [LOGLEVEL][CUSTOMLABEL] Log message.
```

## Import
Run `go get github.com/kenellorando/clog` and add `"github.com/kenellorando/clog"` to your import statements.

## Usage
clog provides five levels of message types (`1-Debug / 2-Info / 3-Warn / 4-Error / 5-Fatal`). All messages will be printed by default, but you may optionally set verbosity with `clog.Level(n)` where `n` is in range [0, 5]. When verbosity is set, only messages at that level and up ("more severe") will be printed.

## Example Application
*main.go*
```Go
package main

import (
	"os"

	"github.com/kenellorando/clog"
)

func example() error {
	clog.Debug("example", "Starting example.")
	clog.Info("example123", "Something is happening!")
	clog.Warn("exampleABC", "Something suspect is happening...")

	_, err := os.Open("fake.txt")
	if err != nil {
		clog.Error("example", "An error has occured. You can choose to pass the error...", err)
		clog.Error("example", "...or pass nil.", nil)

		return err
	}

	return nil
}

func main() {
	clog.Level(5)
	clog.Debug("main", "Staring main.")

	err := example()
	if err != nil {
		clog.Fatal("main", "Fatal triggers an application exit.", nil)
	}

	clog.Debug("main", "Program finished with no errors.")
}
```

*stdout*
```
2019/02/21 12:49:44 [DEBUG][MAIN] Staring main.
2019/02/21 12:49:44 [DEBUG][EXAMPLE] Starting example.
2019/02/21 12:49:44 [INFO ][EXAMPLE123] Something is happening!
2019/02/21 12:49:44 [WARN ][EXAMPLEABC] Something suspect is happening...
2019/02/21 12:49:44 [ERROR][EXAMPLE] An error has occured. You can choose to pass the error...
open fake.txt: no such file or directory
2019/02/21 12:49:44 [ERROR][EXAMPLE] ...or pass nil.
2019/02/21 12:49:44 [FATAL][MAIN] Fatal triggers an application exit.
exit status 1
```
