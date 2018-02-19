package clock

type BallClock struct {
    Min, FiveMin, Hour, Main []int
}

func (c *BallClock) Init(Balls int) {
    c.Min     = []int{}
    c.FiveMin = []int{}
    c.Hour    = []int{}
    c.Main    = []int{}
}

func (c *BallClock) PopBall() int{
    // get the first ball from Main
    // remove that ball from Main
    // return the ball
    return 1
}

func (c *BallClock) AddMinute(ball int) {
    // if the Min queue is 4
        // drop 4 balls back into queue in reverse order
        // clear Min queue
        // append the balls from Min onto Main in reverse order
        // add a ball to FiveMin
    // else 
        // add a new ball to Min
}

func (c *BallClock) AddFiveMin(ball int) {
    // if FiveMin has 11 balls
        // drop balls back in q
        // clear the FiveMin q
        // append the balls from FiveMin onto Main in reverse order
        // add a ball to Hour
    // else 
        // add a new ball to FiveMin
}

func (c *BallClock) AddHour(ball int) {
    // if Hour has 11 balls
        // drop balls back in q
        // clear the Hour q
        // append the balls from FiveMin onto Main in reverse order
        // add a ball to Hour
    // else 
        // add a new ball to FiveMin
}

func (c *BallClock) StepOneMinute() {
    // Get a ball from Main
    // Pass that ball into Add Min
}