# clog
a super simple go logging library

### Usage
`clog` prints log messages to stdout.

**Import**: To import `clog`, run `go get github.com/kenellorando/clog` and add `"github.com/kenellorando/clog"` to your import statements.

**Set Log Level**: To set logging verbosity, you may use `clog.Level(n)` where `n` is a valid integer in range [0, 5]. Running `Level` is optional; `clog` will default to maximum verbosity (5, debug), if no value is received.

`clog`'s logging levels are:
```
0. (disabled)
1. Fatal
2. Error
3. Warn
4. Info
5. Debug
```


### Example
*main.go*
```Go
package main

import (
	"os"

	"github.com/kenellorando/clog"
)

func example() error {
	clog.Debug("example", "Starting example.")
	clog.Info("example", "Something is happening!")
	clog.Warn("example", "Something suspect is happening...")

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
2019/02/21 12:49:44 [INFO ][EXAMPLE] Something is happening!
2019/02/21 12:49:44 [WARN ][EXAMPLE] Something suspect is happening...
2019/02/21 12:49:44 [ERROR][EXAMPLE] An error has occured. You can choose to pass the error...
open fake.txt: no such file or directory
2019/02/21 12:49:44 [ERROR][EXAMPLE] ...or pass nil.
2019/02/21 12:49:44 [FATAL][MAIN] Fatal triggers an application exit.
exit status 1
```
