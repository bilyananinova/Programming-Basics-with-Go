package main

import "fmt"

func main() {
	var students int
	var grade, groupOne, groupTwo, groupThree, groupFour, totalGrades float32

	fmt.Scanln(&students)

	for i := 0; i < students; i++ {
		fmt.Scanln(&grade)

		if grade >= 2.00 && grade <= 2.99 {
			groupFour++
		} else if grade >= 3.00 && grade <= 3.99 {
			groupThree++
		} else if grade >= 4.00 && grade <= 4.99 {
			groupTwo++
		} else if grade >= 5 {
			groupOne++
		}

		totalGrades += grade
	}

	percentGroupOne := groupOne / float32(students) * 100
	percentGroupTwo := groupTwo / float32(students) * 100
	percentGroupThree := groupThree / float32(students) * 100
	percentGroupFour := groupFour / float32(students) * 100
	average := totalGrades / float32(students)

	fmt.Printf("Top students: %.2f%%\n", percentGroupOne)
	fmt.Printf("Between 4.00 and 4.99: %.2f%%\n", percentGroupTwo)
	fmt.Printf("Between 3.00 and 3.99: %.2f%%\n", percentGroupThree)
	fmt.Printf("Fail: %.2f%%\n", percentGroupFour)
	fmt.Printf("Average: %.2f", average)
}
