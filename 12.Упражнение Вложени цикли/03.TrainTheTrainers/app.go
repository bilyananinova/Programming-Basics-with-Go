package main

import "fmt"

func main() {
	var jury, gradesCounter int
	var presentation string
	var totalGradesSum float32

	fmt.Scanln(&jury)
	fmt.Scanln(&presentation)

	for presentation != "Finish" {
		var grade, gradesSum float32

		for i := 0; i < jury; i++ {
			fmt.Scanln(&grade)
			gradesSum += grade
			gradesCounter++

		}
		totalGradesSum += gradesSum

		fmt.Printf("%s - %.2f.\n", presentation, gradesSum/float32(jury))
		fmt.Scanln(&presentation)
	}

	fmt.Printf("Student's final assessment is %.2f.", totalGradesSum/float32(gradesCounter))
}
