package main

import "fmt"

func main() {
	var groupsCount, peopleInGroup, musalla, montBlanc, kilimanjaro, k2, everest, totalPeople int

	fmt.Scanln(&groupsCount)

	for i := 0; i < groupsCount; i++ {

		fmt.Scanln(&peopleInGroup)

		if peopleInGroup <= 5 {
			musalla += peopleInGroup
		} else if peopleInGroup >= 6 && peopleInGroup <= 12 {
			montBlanc += peopleInGroup
		} else if peopleInGroup >= 13 && peopleInGroup <= 25 {
			kilimanjaro += peopleInGroup
		} else if peopleInGroup >= 26 && peopleInGroup <= 40 {
			k2 += peopleInGroup
		} else if peopleInGroup >= 41 {
			everest += peopleInGroup
		}

		totalPeople += peopleInGroup
	}

	fmt.Printf("%.2f%%\n", float32(musalla)/float32(totalPeople)*100)
	fmt.Printf("%.2f%%\n", float32(montBlanc)/float32(totalPeople)*100)
	fmt.Printf("%.2f%%\n", float32(kilimanjaro)/float32(totalPeople)*100)
	fmt.Printf("%.2f%%\n", float32(k2)/float32(totalPeople)*100)
	fmt.Printf("%.2f%%\n", float32(everest)/float32(totalPeople)*100)
}
