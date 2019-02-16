# clog
a super simple go logging library

*Example usage*
```
...

import (
  ...
  "github.com/kenellorando/clog"
)

func example() {
  clog.Debug("example", "Starting example.")
  clog.Info("example", "Something is happening!")
  clog.Warn("example", "Something suspect is happening...")
  
  _, err := os.Open("fake.txt")
  if err != nil {
    clog.Error("example", "An error has occured.", err)
  }
}

func main() {
  clog.Debug("main", "Starting main.")
  example()
  
  _, err := os.Open("fake.txt")
  if err != nil {
    clog.Error("example", "Could not open file for reading!", err)
  }
  clog.Debug("main", "End of program.")
}
```

*Example output*
```
2019/02/15 19:07:50 [DEBUG][MAIN] Staring main.
2019/02/15 19:07:50 [DEBUG][EXAMPLE] Starting example.
2019/02/15 19:07:50 [INFO ][EXAMPLE] Something is happening!
2019/02/15 19:07:50 [WARN ][EXAMPLE] Something suspect is happening...
2019/02/15 19:07:50 [ERROR][EXAMPLE] An error has occured.
open fake.txt: no such file or directory
2019/02/15 19:07:50 [ERROR][MAIN] Could not open file for reading!
open fake.txt: no such file or directory
2019/02/15 19:07:50 [DEBUG][MAIN] End of program.
```
