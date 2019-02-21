# clog
a super simple go logging library

*Example usage*
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
	clog.Init(5)
	clog.Debug("main", "Staring main.")

	err := example()
	if err != nil {
		clog.Fatal("main", "Fatal triggers an application exit.", nil)
	}

	clog.Debug("main", "Program finished with no errors.")
}
```

*Example output*
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
