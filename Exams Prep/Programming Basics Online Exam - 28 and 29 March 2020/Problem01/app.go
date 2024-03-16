package main

import "fmt"

func main() {
	var rent float32
	fmt.Scanln(&rent)

	cake := rent * 0.20
	drinks := cake * 0.55
	animator := rent / 3

	fmt.Println(rent + cake + drinks + animator)
}
