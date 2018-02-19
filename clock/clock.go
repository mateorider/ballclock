package clock

type BallClock struct {
    Min, FiveMin, Hour, Main []int
}

func (c *BallClock) Init(iBalls int) {
    // initial state
    c.Min     = []int{}
    c.FiveMin = []int{}
    c.Hour    = []int{}
    c.Main    = []int{}

    // populate the Main slice
    for i := 1; i <= iBalls; i++ {
        c.Main = append(c.Main, i)
    }
}

func (c *BallClock) PopBall() int{
    // get the first ball from Main
    firstBall := c.Main[0]
    // remove that ball from Main
    c.Main = removeIndex(c.Main, 0)
    return firstBall
}

func (c *BallClock) AddMinute(iBall int) {
    if len(c.Min) == 4 {
        // reverse the order of the Min balls
        reversedMin := reverseSlice(c.Min)
        // clear Min queue
        c.Min = []int{}
        // append the reversed Min balls to Main
        c.Main = append(c.Main, reversedMin...)
        // add a ball to FiveMin
        c.AddFiveMin(iBall)
    } else {
        // add a new ball to Min
        c.Min = append(c.Min, iBall)
    }
}

func (c *BallClock) AddFiveMin(iBall int) {
    if len(c.FiveMin) == 11 {
        // drop balls back in q
        reversedFiveMin := reverseSlice(c.FiveMin)
        // clear the FiveMin q
        c.FiveMin = []int{}
        // append the balls from FiveMin onto Main in reverse order
        c.Main = append(c.Main, reversedFiveMin...)
        // add a ball to Hour
        c.AddHour(iBall)
    } else {
        // add a new ball to FiveMin
        c.FiveMin = append(c.FiveMin, iBall)
    }
}

func (c *BallClock) AddHour(iBall int) {
    if len(c.Hour) == 11 {
        // append the balls to Main first on this this time
        reversedHour := reverseSlice(c.Hour)
        c.Main = append(c.Main, reversedHour...)
        c.Main = append(c.Main, iBall)
        // clear the Hour q last
        c.Hour = []int{}
    } else {
        c.Hour = append(c.Hour, iBall)
    }
}

func (c *BallClock) StepOneMinute() {
    // Get a ball from Main
    nextBall := c.PopBall()
    // Pass that ball into Add Min
    c.AddMinute(nextBall)
}

// --------------------------------
//  UTILS
// --------------------------------

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