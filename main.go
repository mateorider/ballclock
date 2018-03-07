package main

import (
    "flag"
    "fmt"
    "log"
    "github.com/rider-mateo/ballclock/clock"
    "reflect"
    "time"
)

// Mode1 takes a single parameter specifying the number of balls and reports the number of balls 
// given in the input and the number of days (24-hour periods) which elapse before the clock returns
// to its initial ordering. 
// Returns days int, seconds float64
func Mode1(ballCount int) (int, float64) {
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

// validateFlagInput takes the name of a flag, the value from the flag, the lower bound, and the 
// upper bound to check if the flag input falls within the valid range for the program. Logs the
// error message and exits the program if the value is outside of the acceptable range.
func validateFlagInput(name string, count, lower, upper int) {
    if count < lower || count > upper {
        log.Fatalf("%s must be a number between %d and %d", name, lower, upper)
    }
    return
}

func main() {
    mode    := flag.Int("mode", 1, "The mode the clock will operate in where: 1 = cycledays, 2 = clock state")
    balls   := flag.Int("balls", 27, "The number of balls to be used in the clock. Valid numbers between 27 and 127.")
    minutes := flag.Int("minutes", 720, "Duration to run in minutes, Valid numbers between 720 and 1440. MODE 2 ONLY")

    flag.Parse()
    
    validateFlagInput("mode", *mode, 1, 2)
    validateFlagInput("balls", *balls, 27, 127)
    validateFlagInput("minutes", *minutes, 720, 1440)

    if *mode == 1 {

        days, seconds := Mode1(*balls)

        fmt.Printf("%v balls cycle after %v days.\n", *balls, days)
        fmt.Printf("Completed in %0.1f miliseconds (%0.3f seconds)", seconds*1000, seconds)

    } else {

        clock := Mode2(*balls, *minutes)

        fmt.Println(clock.ToString())
    
    } 
}