package main

import "fmt"

func main() {
	var students, a, b, c, d, total int
	var grade, average float32

	fmt.Scanln(&students)

	for i := 0; i < students; i++ {
		fmt.Scanln(&grade)

		if grade >= 5 {
			a++
		} else if grade >= 4 && grade <= 4.99 {
			b++
		} else if grade >= 3 && grade <= 3.99 {
			c++
		} else if grade < 3 {
			d++
		}

		average += grade
	}

	total = a + b + c + d
	aPercent := float32(a) / float32(total) * 100
	bPercent := float32(b) / float32(total) * 100
	cPercent := float32(c) / float32(total) * 100
	dPercent := float32(d) / float32(total) * 100
	average = average / float32(total)

	fmt.Printf("Top students: %.2f%%\n", aPercent)
	fmt.Printf("Between 4.00 and 4.99: %.2f%%\n", bPercent)
	fmt.Printf("Between 3.00 and 3.99: %.2f%%\n", cPercent)
	fmt.Printf("Fail: %.2f%%\n", dPercent)
	fmt.Printf("Average: %.2f\n", average)

}
