package main

import "fmt"

func main() {
	var oneLv, twoLv, fiveLv, sum int

	fmt.Scan(&oneLv, &twoLv, &fiveLv, &sum)

	for i := 0; i <= oneLv; i++ {
		for j := 0; j <= twoLv; j++ {
			for k := 0; k <= fiveLv; k++ {

				if (i*1)+(j*2)+(k*5) == sum {
					fmt.Printf("%d * 1 lv. + %d * 2 lv. + %d * 5 lv. = %d lv.\n", i, j, k, (i*1)+(j*2)+(k*5))
				}
			}
		}
	}
}
