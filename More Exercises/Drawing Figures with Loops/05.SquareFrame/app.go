package main

import "fmt"

func main() {
	var n int
	var str string

	fmt.Scanln(&n)

	for row := 1; row <= n; row++ {
		str = "+ "

		for col := 1; col <= n-2; col++ {
			str += "- "
		}

		str += "+"
		fmt.Println(str)

		for m := 0; m < n-2; m++ {
			str = "| "

			for col := 0; col < n-2; col++ {
				str += "- "
			}

			str += "|"
			fmt.Println(str)
			continue
		}

		for e := 1; e < n; e++ {
			str = "+ "

			for col := 1; col <= n-2; col++ {
				str += "- "
			}

			str += "+"
			fmt.Println(str)
			break
		}
		break
	}
}
