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
func Mode1 (ballCount int) (int, float64) {
    start := time.Now()
    days := 0

    // create a two instances of BallClock, one to iterate and one to compare to
    var ballClock, initialOrdering clock.BallClock
    ballClock.Init(ballCount)
    initialOrdering.Init(ballCount)

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

// Mode2 takes two parameters, the number of balls and the number of minutes to run for. If the
// number of minutes is specified, the clock must run to the number of minutes and report the state 
// of the tracks at that point in a JSON format.
func Mode2(ballCount int, minuteCount int) clock.BallClock {
    
    // create one new instance of BallClock
    var ballClock clock.BallClock
    ballClock.Init(ballCount)

    // Iterate until minuteCount is reached
    for i := 1; i <= minuteCount; i++ {
        ballClock.StepOneMinute()
    }

    return ballClock
}