package main

import "fmt"

func main() {
	var x, y, h float32

	fmt.Scan(&x, &y, &h)

	frontAndBackWalls := 2*(x*x) - (1.2 * 2)
	sideWalls := 2*(x*y) - 2*(1.5*1.5)
	walls := frontAndBackWalls + sideWalls
	littersGreenPaint := walls / 3.4

	frontAndBackRoof := 2 * (x * h / 2)
	sideRoof := 2 * (x * y)
	roof := frontAndBackRoof + sideRoof
	littersRedPaint := roof / 4.3

	fmt.Printf("%.2f\n", littersGreenPaint)
	fmt.Printf("%.2f", littersRedPaint)

}
