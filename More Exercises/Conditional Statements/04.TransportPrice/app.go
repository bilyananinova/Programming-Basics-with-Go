package main

import "fmt"

func main() {
	var kms int
	var time string
	var price float32

	fmt.Scan(&kms, &time)

	if kms >= 100 {

		price = 0.06 * float32(kms)

	} else if kms >= 20 {

		price = 0.09 * float32(kms)

	} else {

		if time == "day" {
			price = 0.70 + float32(kms)*0.79
		} else if time == "night" {
			price = 0.70 + float32(kms)*0.90
		}

	}

	fmt.Printf("%.2f", price)
}
