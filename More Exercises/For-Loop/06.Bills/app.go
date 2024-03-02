package main

import "fmt"

func main() {
	var months int
	var electricity, water, internet, other, sum float64

	fmt.Scanln(&months)

	for i := 1; i <= months; i++ {
		fmt.Scanln(&sum)

		electricity += sum
		water += 20
		internet += 15
	}

	other = (electricity + water + internet) * 1.2

	bills := electricity + water + internet + other
	average := bills / float64(months)

	fmt.Printf("Electricity: %.2f lv\n", electricity)
	fmt.Printf("Water: %.2f lv\n", water)
	fmt.Printf("Internet: %.2f lv\n", internet)
	fmt.Printf("Other: %.2f lv\n", other)
	fmt.Printf("Average: %.2f lv", average)
}
