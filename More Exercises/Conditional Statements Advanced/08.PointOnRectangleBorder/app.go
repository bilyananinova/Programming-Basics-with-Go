package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x, y float32

	fmt.Scan(&x1, &y1, &x2, &y2, &x, &y)

	if (x == x1 || x == x2) && (y1 <= y && y <= y2) {
		fmt.Println("Border")
	} else if (y == y1 || y == y2) && (x1 <= x && x <= x2) {
		fmt.Println("Border")
	} else {
		fmt.Println("Inside / Outside")
	}

}
