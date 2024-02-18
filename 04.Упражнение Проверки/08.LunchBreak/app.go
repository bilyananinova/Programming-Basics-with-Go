package main

import (
	"fmt"
	"math"
)

func main() {
	var serial string
	var episodeTime, breakTimeLong int

	fmt.Scanln(&serial)
	fmt.Scanln(&episodeTime)
	fmt.Scanln(&breakTimeLong)

	lunchTime := float64(breakTimeLong) / 8.0
	restTime := float64(breakTimeLong) / 4.0
	leftTime := float64(breakTimeLong) - lunchTime - restTime

	if leftTime >= float64(episodeTime) {
		fmt.Printf("You have enough time to watch %s and left with %.f minutes free time.", serial, math.Ceil(leftTime-float64(episodeTime)))
	} else {
		fmt.Printf("You don't have enough time to watch %s, you need %.f more minutes.", serial, math.Ceil(float64(episodeTime)-leftTime))
	}
}
