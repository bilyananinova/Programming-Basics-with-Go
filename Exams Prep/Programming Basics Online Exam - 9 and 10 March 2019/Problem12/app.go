package main

import "fmt"

func main() {
	var limitTimeInMinutes, limitTimeInSeconds, secondForHundredMeters int
	var lengthOfChuteInMeters float64

	fmt.Scanln(&limitTimeInMinutes)
	fmt.Scanln(&limitTimeInSeconds)
	fmt.Scanln(&lengthOfChuteInMeters)
	fmt.Scanln(&secondForHundredMeters)

	totalLimitTimeInSeconds := float64(limitTimeInMinutes*60) + float64(limitTimeInSeconds)
	delayTime := (lengthOfChuteInMeters / 120) * 2.5
	playerTime := (lengthOfChuteInMeters/100)*float64(secondForHundredMeters) - delayTime

	if totalLimitTimeInSeconds >= float64(playerTime) {
		fmt.Println("Marin Bangiev won an Olympic quota!")
		fmt.Printf("His time is %.3f.", playerTime)
	} else {
		fmt.Printf("No, Marin failed! He was %.3f second slower.", playerTime-totalLimitTimeInSeconds)
	}
}
