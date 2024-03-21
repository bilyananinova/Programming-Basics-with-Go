package main

import "fmt"

func main() {
	var rent int
	fmt.Scanln(&rent)

	statuettes := float64(rent) * 0.70
	catering := statuettes * 0.85
	voiceover := catering / 2

	total := float64(rent) + statuettes + catering + voiceover

	fmt.Printf("%.2f", total)

}
