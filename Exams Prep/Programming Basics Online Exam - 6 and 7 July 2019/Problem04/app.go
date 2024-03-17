package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var target, income float32
	var coctail string
	var order int
	fmt.Scanln(&target)
	scanner.Scan()
	coctail = scanner.Text()

	for coctail != "Party!" {
		fmt.Scanln(&order)

		var price float32 = float32(len(coctail) * order)

		if int(price)%2 == 1 {
			price = float32(price * 0.75)
		}

		income += float32(price)

		if income >= target {
			break
		}

		scanner.Scan()
		coctail = scanner.Text()
	}

	if coctail == "Party!" {
		fmt.Printf("We need %.2f leva more.\n", target-income)
	} else {
		fmt.Println("Target acquired.")
	}

	fmt.Printf("Club income - %.2f leva.", income)

}

