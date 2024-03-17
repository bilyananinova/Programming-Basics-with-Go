package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var profit, money float32
	var coctail string
	var order int
	fmt.Scanln(&profit)
	fmt.Scanln(&coctail)

	for coctail != "Party!" {
		fmt.Scanln(&order)

		price := float32(len(coctail) * order)

		if int(price)%2 == 1 {
			price *= 0.75
		}

		money += price

		if money >= profit {
			fmt.Println("Target acquired.")
			break
		}

		if coctail == "Party!" {
			break
		}

		scanner.Scan()
		coctail = scanner.Text()
	}

	if coctail == "Party!" {
		fmt.Printf("We need %.2f leva more.\n", profit-money)
	}

	fmt.Printf("Club income - %.2f leva.", money)

}
