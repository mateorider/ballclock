package clock

import "encoding/json"

// BallClock represents a ball clock with four trays (slices): 
// (1) Min holds balls that represent individual minutes passed
// (2) FiveMin holds balls that represent 5-minutes periods
// (3) Hour holds balls for the # of hours passed
// (4) Main holds any unused balls 
type BallClock struct {
    Min, FiveMin, Hour, Main []int
}

// Init initializes the state of the ballclock, taking an integer that represents the number of
// balls to be used in the clock.
func (c *BallClock) Init(balls int) {
    // initial state
    c.Min     = []int{}
    c.FiveMin = []int{}
    c.Hour    = []int{}
    c.Main    = []int{}

    // populate the Main slice
    for i := 1; i <= balls; i++ {
        c.Main = append(c.Main, i)
    }
}

// PopBall removes and returns the first ball from the Main slice
func (c *BallClock) PopBall() int{
    firstBall := c.Main[0]
    c.Main = removeIndex(c.Main, 0)
    return firstBall
}

// AddMinute adds a ball to the Min slice. If the slice is full, the contents of the slice are
// reversed, emptied from Min and appended to the Main slice while a single ball is added to the 
// FiveMin slice.
func (c *BallClock) AddMinute(ball int) {
    if len(c.Min) == 4 {
        reversedMin := reverseSlice(c.Min)
        c.Min = []int{}
        c.Main = append(c.Main, reversedMin...)
        c.AddFiveMin(ball)
    } else {
        c.Min = append(c.Min, ball)
    }
}

// AddFive adds a ball to the FiveMin slice. If the slice is full, the contents of the slice are
// reversed, emptied from FiveMin and appended to the Main slice while a single ball is added to the 
// Hour slice.
func (c *BallClock) AddFiveMin(ball int) {
    if len(c.FiveMin) == 11 {
        reversedFiveMin := reverseSlice(c.FiveMin)
        c.FiveMin = []int{}
        c.Main = append(c.Main, reversedFiveMin...)
        c.AddHour(ball)
    } else {
        c.FiveMin = append(c.FiveMin, ball)
    }
}

// AddHour adds a ball to the Hour slice. If the slice is full, all slices must be emptied into 
// the Main slice to reset the clock. The reversed Hour slice is appeneded to Main, followed by the
// added ball. 
func (c *BallClock) AddHour(ball int) {
    if len(c.Hour) == 11 {
        reversedHour := reverseSlice(c.Hour)
        c.Main = append(c.Main, reversedHour...)
        c.Main = append(c.Main, ball)
        c.Hour = []int{}
    } else {
        c.Hour = append(c.Hour, ball)
    }
}

// StepOneMinute retrieves a ball from the Main slice and adds it to the Min slice.
func (c *BallClock) StepOneMinute() {
    nextBall := c.PopBall()
    c.AddMinute(nextBall)
}

// ToString returns a string representation of the BallClock's state (ordering) in JSON format. 
func (c *BallClock) ToString() string{
    jsonBytes, err := json.Marshal(c)

    if err != nil {
        panic(err)
    }

    return string(jsonBytes)
}

// reverseSlice takes a slice and returns a reverse copy of it
func reverseSlice(original []int) []int {
    rev := make([]int, len(original))
    
    copy(rev, original)

    for beg, end := 0, len(rev)-1; beg < end; beg, end = beg+1, end-1 {
        rev[beg], rev[end] = rev[end], rev[beg]
    }
    return rev
}

// removeIndex takes a slice and an index, and returns a copy of the slice
// without the specified index
func removeIndex(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}