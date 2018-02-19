package main

import (
    "flag"
    "fmt"
    "log"
    "github.com/rider-mateo/ballclock/modes"
)

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

        iDays, fSeconds := modes.Mode1(*balls)

        fmt.Printf("%v balls cycle after %v days.\n", *balls, iDays)
        fmt.Printf("Completed in %0.1f miliseconds (%0.3f seconds)", fSeconds*1000, fSeconds)

    } else {

        clock := modes.Mode2(*balls, *minutes)

        fmt.Println(clock.ToString())
    
    } 
}