package modes

import (
    "reflect"
    "time"
    "github.com/rider-mateo/ballclock/clock"
)

// Mode1 takes a single parameter specifying the number of balls and reports the number of balls 
// given in the input and the number of days (24-hour periods) which elapse before the clock returns
// to its initial ordering. 
// Returns days int, seconds float64
func Mode1 (iBalls int) (int, float64) {
    start := time.Now()
    days := 0

    // create a two instances of BallClock, one to iterate and one to compare to
    var ballClock, initialOrdering clock.BallClock
    ballClock.Init(iBalls)
    initialOrdering.Init(iBalls)

    // Offset the states
    ballClock.StepOneMinute()
    minutes := 1

    // Until ballClock matches the initialClockState, add minutes
    for !reflect.DeepEqual(ballClock, initialOrdering) {
        ballClock.StepOneMinute()
        minutes++
    }
    
    // stop the timer and calculate days
    seconds := time.Since(start).Seconds()
    days = minutes / (60 * 24)

    return days, seconds
}