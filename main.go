package main

import (
    "flag"
    "fmt"
    "log"
    "github.com/rider-mateo/ballclock/modes"
)

func validateBallInput(count int) {
    if count < 27 || count > 127 {
        log.Fatal("'balls' must be a valid number between 27 and 127")
    }
}

func main() {
    mode    := flag.Int("mode", 1, "The mode the clock will operate in where: 1 = cycledays, 2 = clock state")
    balls   := flag.Int("balls", 27, "The number of balls to be used in the clock. Valid numbers between 27 and 127.")
    minutes := flag.Int("minutes", 720, "Duration to run in minutes, MODE 2 ONLY, ")

    flag.Parse()

    validateBallInput(*balls)

    if *mode == 1 {

        iDays, fSeconds := modes.Mode1(*balls)

        fmt.Printf("%v balls cycle after %v days.\n", *balls, iDays)
        fmt.Printf("Completed in %0.1f miliseconds (%0.3f seconds)", fSeconds*1000, fSeconds)

    } else if *mode == 2 {

        fmt.Println("mode 2 using", *balls, *minutes)

    } else {

        fmt.Println("not a valid mode")

    }
}