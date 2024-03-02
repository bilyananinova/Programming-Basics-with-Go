package main

import "fmt"

func main() {
	var loads int
	var sum, tons, van, truck, train float32

	fmt.Scanln(&loads)

	for i := 0; i < loads; i++ {
		fmt.Scanln(&tons)

		if tons <= 3 {
			van += tons
		} else if tons >= 4 && tons <= 11 {
			truck += tons
		} else if tons >= 12 {
			train += tons
		}

	}

	sum = ((van * 200) + (truck * 175) + (train * 120)) / (van + truck + train)

	fmt.Printf("%.2f\n", sum)
	fmt.Printf("%.2f%%\n", van/(van+train+truck)*100)
	fmt.Printf("%.2f%%\n", truck/(van+train+truck)*100)
	fmt.Printf("%.2f%%\n", train/(van+train+truck)*100)
}
