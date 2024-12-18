package day8

import (
	"fmt"

	"github.com/TTF5/AoC2024/utility"
)

func generateAntinodes2(radarMap RadarMap, antennas map[string][]vec) map[utility.Point]int {
	antinodeSet := make(map[utility.Point]int)
	for _, positions := range antennas {

		if len(positions) < 2 {
			continue
		}

		pairs := generatePairs(positions)
		for _, pair := range pairs {
			a := pair[0]
			b := pair[1]
			diff := a.Sub(b) // Vector from b to a

			for i := 1; true; i++ {
				antinode := b.Add(diff.Scale(float64(i)))
				if antinode.X() < 0 || antinode.X() >= float64(radarMap.width) {
					break
				}
				if antinode.Y() < 0 || antinode.Y() >= float64(radarMap.height) {
					break
				}
				antinodeSet[utility.Point{X: int(antinode.X()), Y: int(antinode.Y())}]++
			}
			for i := 1; true; i++ {
				antinode := b.Add(diff.Scale(float64(-i)))
				if antinode.X() < 0 || antinode.X() >= float64(radarMap.width) {
					break
				}
				if antinode.Y() < 0 || antinode.Y() >= float64(radarMap.height) {
					break
				}
				antinodeSet[utility.Point{X: int(antinode.X()), Y: int(antinode.Y())}]++
			}
		}
	}
	return antinodeSet
}

func Day8Part2() {
	radarMap := parseInput("day8/input.txt")
	printMap(radarMap)

	antennas := getAntennas(radarMap)
	antinodeMap := generateAntinodes2(radarMap, antennas)

	numAntinodes := len(antinodeMap)

	fmt.Println(numAntinodes)
}
