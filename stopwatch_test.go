package stopwatch

import (
	"fmt"
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {

	myStopwatch := NewStopwatch()
	myStopwatch.Start()
	time.Sleep(100 * time.Millisecond)
	myStopwatch.Stop()

	actualString := myStopwatch.SecondsFormatted(2)
	expectedString := "0.10"
	errorMessage := assert(expectedString, actualString)
	if errorMessage != nil {
		t.Errorf(*errorMessage)
	}
}

func TestStopWatchDelayedCheck(t *testing.T) {

	myStopwatch := NewStopwatch()
	myStopwatch.Start()
	time.Sleep(100 * time.Millisecond)
	myStopwatch.Stop()
	time.Sleep(100 * time.Millisecond)

	actualString := myStopwatch.SecondsFormatted(2)
	expectedString := "0.10"
	errorMessage := assert(expectedString, actualString)
	if errorMessage != nil {
		t.Errorf(*errorMessage)
	}
}

func TestStopWatchStopAndRestart(t *testing.T) {

	myStopwatch := NewStopwatch()
	myStopwatch.Start()
	time.Sleep(100 * time.Millisecond)

	actualString := myStopwatch.SecondsFormatted(2)
	expectedString := "0.10"
	errorMessage := assert(expectedString, actualString)
	if errorMessage != nil {
		t.Errorf(*errorMessage)
	}

	time.Sleep(100 * time.Millisecond)
	if myStopwatch.IsStopped() {
		t.Errorf("StopWatch was never stopped. The internal state must be corrupted.")
	}
	if myStopwatch.Start() {
		t.Errorf("StopWatch was never stopped. calling Start should return false.")
	}
	actualString = myStopwatch.SecondsFormatted(2)
	expectedString = "0.20"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}

	myStopwatch.Reset()
	time.Sleep(100 * time.Millisecond)
	actualString = myStopwatch.SecondsFormatted(2)
	expectedString = "0.10"
	errorMessage = assert(expectedString, actualString)
	if errorMessage != nil {
		t.Errorf(*errorMessage)
	}

}

func assert(expected string, actual string) *string {
	if expected != actual {
		errorMessage := fmt.Sprintf("Expected String(%s) is not same as"+
			" actual string (%s).", expected, actual)
		return &errorMessage
	}
	return nil
}
