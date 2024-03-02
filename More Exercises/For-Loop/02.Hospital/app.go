package main

import "fmt"

func main() {
	var days, patients, treated, untreated int
	var doctors int = 7

	fmt.Scanln(&days)

	for i := 1; i <= days; i++ {
		fmt.Scanln(&patients)

		if i%3 == 0 {
			if untreated > treated {
				doctors++
			}
		}

		if patients > doctors {
			treated += doctors
			untreated += patients - doctors
		} else {
			treated += patients
		}

	}

	fmt.Printf("Treated patients: %d.\n", treated)
	fmt.Printf("Untreated patients: %d.", untreated)
}
