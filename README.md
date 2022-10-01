# stopwatch
A simple go utility for creating a stop watch.  The current time can be read even if it is not stopped yet.  It can be resumed again after stopping.

## Example

```go
package stopwatch

import (
	"time"
	"github.com/go-youtils/stopwatch"
)

func main() {

	myStopwatch := NewStopwatch()
	myStopwatch.Start()

	time.Sleep(100 * time.Millisecond)

	currentTimeInSeconds := myStopwatch.SecondsFormatted(2)
	println("Current time: " + currentTimeInSeconds) //should be equal to "0.10"

	time.Sleep(100 * time.Millisecond)

	if myStopwatch.IsStopped() {
		//We should not enter this block since the stopwatch is not stopped
		myStopwatch.Start()
	}
	if myStopwatch.Start() {
		println("I should not be able to start an already started stopwatch")
	}

	currentTimeInSeconds = myStopwatch.SecondsFormatted(2)
	println("Current time: " + currentTimeInSeconds) //should be equal to "0.20"

	myStopwatch.Reset()

	time.Sleep(100 * time.Millisecond)

	currentTimeInSeconds = myStopwatch.SecondsFormatted(2)
	println("Current time: " + currentTimeInSeconds) //should be equal to "0.10"
}
```

