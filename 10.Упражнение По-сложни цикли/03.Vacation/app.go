package main

import "fmt"

func main() {
	var neededMoney, currentMoney, sum float32
	var action string
	var spendDays, days int

	fmt.Scanln(&neededMoney)
	fmt.Scanln(&currentMoney)

	for currentMoney < neededMoney {
		fmt.Scanln(&action)
		fmt.Scanln(&sum)

		days++

		if action == "spend" {
			spendDays++
			currentMoney -= sum

			if currentMoney < sum {
				currentMoney = 0
			}

			if spendDays == 5 {
				fmt.Println("You can't save the money.")
				fmt.Println(days)
				break
			}
		}

		if action == "save" {
			currentMoney += sum
			spendDays = 0

			if currentMoney >= neededMoney {
				fmt.Printf("You saved the money for %d days.", days)
				break
			}
		}
	}

}
