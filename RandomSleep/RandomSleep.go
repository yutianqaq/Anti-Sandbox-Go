package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomSleep() int {
	startTime := time.Now()
	randomInt := rand.Intn(100)
	fmt.Print(randomInt)
	randomDuration := time.Duration(randomInt) * time.Second
	time.Sleep(randomDuration)
	endTime := time.Now()
	sleepTime := endTime.Sub(startTime)
	if sleepTime >= randomDuration {
		return 1
	} else {
		return 0
	}
}

func main() {
	if RandomSleep() == 0 {
		return
	}
}
