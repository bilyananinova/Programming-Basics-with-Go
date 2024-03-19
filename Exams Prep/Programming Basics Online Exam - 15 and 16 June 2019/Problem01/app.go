package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var serial string
	var seasons, episodes int
	var episodeTimeWithoutAdv float32

	scanner.Scan()
	serial = scanner.Text()

	// fmt.Scanln(&serial)
	fmt.Scanln(&seasons)
	fmt.Scanln(&episodes)
	fmt.Scanln(&episodeTimeWithoutAdv)

	episodeTimeWithAdv := float32(episodeTimeWithoutAdv) * 1.20
	additionalTimeForLastEpisodes := seasons * 10
	totalTime := (float32(seasons) * float32(episodes) * episodeTimeWithAdv) + float32(additionalTimeForLastEpisodes)

	fmt.Printf("Total time needed to watch the %s series is %.0f minutes.", serial, totalTime)
}
