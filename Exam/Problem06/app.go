package main

import "fmt"

func main() {
	var firstEnd, secondEnd, thirdEnd, counter int

	fmt.Scanln(&firstEnd)
	fmt.Scanln(&secondEnd)
	fmt.Scanln(&thirdEnd)

	for first := 1; first <= firstEnd; first++ {
		for second := 1; second <= secondEnd; second++ {
			for third := 1; third <= thirdEnd; third++ {
				if first%2 == 0 && third%2 == 0 {

					for i := 1; i <= second; i++ {
						if second%i == 0 {
							counter++
						}
					}

					if counter == 2 {
						fmt.Printf("%d %d %d\n", first, second, third)
					}

					counter = 0
				}
			}
		}
	}
}
