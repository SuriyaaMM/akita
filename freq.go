package core

import (
	"math"
)

// Freq defines the type of frequency
type Freq float64

// Defines the unit of frequency
const (
	Hz  Freq = 1
	KHz Freq = 1e3
	MHz Freq = 1e6
	GHz Freq = 1e9
)

// Period returns the time between two consecutive ticks
func (f Freq) Period() VTimeInSec {
	return VTimeInSec(1.0 / f)
}

// NextTick returns the next tick time.
//
// If currTime is not on a tick time, this function returns the time of
// upcomming tick.
func (f Freq) NextTick(currTime VTimeInSec) VTimeInSec {
	period := f.Period()
	count := math.Floor(float64((currTime + period*1e-6) / period))
	return VTimeInSec(count+1) * period
}

// NCyclesLater returns the time after N cycles
//
// This function will always return a time of an integer number of cycles
func (f Freq) NCyclesLater(n int, currTime VTimeInSec) VTimeInSec {
	return f.NextTick(currTime + VTimeInSec(n)*f.Period())
}

// NoEarlierThan returns the tick time that is at or right after the given time
func (f Freq) NoEarlierThan(t VTimeInSec) VTimeInSec {
	count := t / f.Period()
	return VTimeInSec(math.Ceil(float64(count))) * f.Period()
}
