package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var actorsName, juryName string
	var academyPoints, juryPoints float64
	var jury int

	scanner.Scan()
	actorsName = scanner.Text()

	fmt.Scanln(&academyPoints)
	fmt.Scanln(&jury)

	for i := 0; i < jury; i++ {
		scanner.Scan()
		juryName = scanner.Text()
		fmt.Scanln(&juryPoints)

		academyPoints += (float64(len(juryName)) * float64(juryPoints)) / 2

		if academyPoints >= 1250.5 {
			break
		}
	}

	if academyPoints >= 1250.5 {
		fmt.Printf("Congratulations, %s got a nominee for leading role with %.1f!", actorsName, academyPoints)
	} else {
		fmt.Printf("Sorry, %s you need %.1f more!", actorsName, 1250.5-academyPoints)
	}
}
