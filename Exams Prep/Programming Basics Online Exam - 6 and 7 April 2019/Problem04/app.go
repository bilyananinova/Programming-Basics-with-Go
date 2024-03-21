package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var voucher, tickets, others int
	var input string

	fmt.Scanln(&voucher)

	scanner.Scan()
	input = scanner.Text()

	for input != "End" {

		if len(input) > 8 {

			if int(input[0])+int(input[1]) > voucher {
				break
			}

			tickets++

			voucher -= int(input[0]) + int(input[1])
		} else {

			if int(input[0]) > voucher {
				break
			}

			others++

			voucher -= int(input[0])
		}

		scanner.Scan()
		input = scanner.Text()
	}

	fmt.Println(tickets)
	fmt.Println(others)
}
