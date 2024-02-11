package main

import "fmt"

func main() {
	var bgn float32 = 1.79549
	var usd float32

	fmt.Scanln(&usd)

	usdToBgn := usd * bgn

	fmt.Println(usdToBgn)
}
