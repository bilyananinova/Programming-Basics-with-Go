package main

import "fmt"

func main() {
	var firstNumberEnd, secondNumberEnd, thirdNumberEnd int

	fmt.Scan(&firstNumberEnd, &secondNumberEnd, &thirdNumberEnd)

	for i := 1; i <= firstNumberEnd; i++ {
		for j := 1; j <= secondNumberEnd; j++ {
			for k := 1; k <= thirdNumberEnd; k++ {

				if i%2 == 0 {

					if j == 2 || j == 3 || j == 5 || j == 7 {

						if k%2 == 0 {
							fmt.Printf("%d %d %d\n", i, j, k)
						}
					}
				}

			}
		}
	}

}
