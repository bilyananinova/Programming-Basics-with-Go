package main

import "fmt"

func main() {
	var length int
	var height int
	var width int
	var percentage float32

	fmt.Scanln(&length)
	fmt.Scanln(&height)
	fmt.Scanln(&width)
	fmt.Scanln(&percentage)

	var volumeAqariumCM float32 = float32(length) * float32(height) * float32(width)
	var volumeAqariumDM float32 = volumeAqariumCM * 0.001
	var content float32 = volumeAqariumDM * (percentage / 100)
	var litters float32 = volumeAqariumDM - content

	fmt.Println(litters)
}
