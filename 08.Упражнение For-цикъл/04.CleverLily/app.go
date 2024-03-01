package main

import "fmt"

func main() {
	var n, birthDayToys int
	var birthDayMoney, money, washingMachinePrice, toysPrice float32

	fmt.Scanln(&n)
	fmt.Scanln(&washingMachinePrice)
	fmt.Scanln(&toysPrice)

	for i := 1; i <= n; i++ {

		if i%2 == 0 {
			birthDayMoney += 5.0*float32(i) - 1
		} else {
			birthDayToys++
		}
	}

	money = birthDayMoney + float32(birthDayToys)*toysPrice

	if money >= washingMachinePrice {
		left := money - washingMachinePrice
		fmt.Printf("Yes! %.2f", left)
	} else {
		needed := washingMachinePrice - money
		fmt.Printf("No! %.2f", needed)
	}

}
