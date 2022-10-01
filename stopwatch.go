package stopwatch

import (
	"fmt"
	"strconv"
	"time"
)

type Stopwatch struct {
	lastStartTime   time.Time
	lastPauseTime   time.Time
	hasStarted      bool
	runningDuration float64
	isStopped       bool
}

func NewStopwatch() *Stopwatch {
	s := new(Stopwatch)
	s.Reset()
	return s
}

func (s *Stopwatch) Start() bool {
	if !s.isStopped {
		return false
	}
	s.hasStarted = true
	s.isStopped = false
	s.lastStartTime = time.Now()
	return true
}

func (s *Stopwatch) Reset() {
	s.lastPauseTime = time.Now()
	s.hasStarted = false
	s.isStopped = true
	s.runningDuration = 0.0
}

func (s *Stopwatch) Stop() {
	s.isStopped = true
	s.lastPauseTime = time.Now()

	s.runningDuration += s.calculateCurrentDuration()
}

func (s *Stopwatch) IsStopped() bool {
	return s.isStopped
}

func (s *Stopwatch) SecondsFormatted(decimals int) string {
	format := "%." + strconv.Itoa(decimals) + "f"
	return fmt.Sprintf(format, s.Seconds())
}

func (s *Stopwatch) Seconds() float64 {

	if !s.hasStarted {
		return 0.0
	}

	if s.isStopped {
		return s.runningDuration
	}

	return s.runningDuration + s.calculateCurrentDuration()
}

func (s *Stopwatch) calculateCurrentDuration() float64 {
	return time.Now().Sub(s.lastStartTime).Seconds()
}
