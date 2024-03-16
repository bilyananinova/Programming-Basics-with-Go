package main

import "fmt"

func main() {
	var groups, group, total int
	var m1, m2, m3, m4, m5 float32
	fmt.Scanln(&groups)

	for i := 1; i <= groups; i++ {
		fmt.Scanln(&group)

		if group <= 5 {
			m1 += float32(group)
		} else if group >= 6 && group <= 12 {
			m2 += float32(group)
		} else if group >= 13 && group <= 25 {
			m3 += float32(group)
		} else if group >= 26 && group <= 40 {
			m4 += float32(group)
		} else if group > 40 {
			m5 += float32(group)
		}

		total += group
	}

	p1 := m1 / float32(total) * 100
	p2 := m2 / float32(total) * 100
	p3 := m3 / float32(total) * 100
	p4 := m4 / float32(total) * 100
	p5 := m5 / float32(total) * 100
	fmt.Printf("%.2f%%\n", p1)
	fmt.Printf("%.2f%%\n", p2)
	fmt.Printf("%.2f%%\n", p3)
	fmt.Printf("%.2f%%\n", p4)
	fmt.Printf("%.2f%%", p5)

}
