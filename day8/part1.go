package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TTF5/AoC2024/utility"
	"github.com/quartercastle/vector"
)

type vec = vector.Vector

type RadarMap struct {
	width    int
	height   int
	data     [][]string
	antennas map[rune]vec
}

func parseInput(input string) RadarMap {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var radarMap RadarMap

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		if radarMap.width == 0 {
			radarMap.width = len(line)
		}
		radarMap.data = append(radarMap.data, row)
	}
	radarMap.height = len(radarMap.data)
	return radarMap
}

func getAntennas(radarMap RadarMap) map[string][]vec {
	antennas := make(map[string][]vec)
	for y, row := range radarMap.data {
		for x, cell := range row {
			if cell != "." && cell != "#" {
				antennas[cell] = append(antennas[cell], vec{float64(x), float64(y)})
			}
		}
	}
	return antennas
}

func generatePairs(points []vec) [][2]vec {
	pairs := make([][2]vec, 0)
	for i, a := range points {
		for j, b := range points {
			if i != j {
				pairs = append(pairs, [2]vec{a, b})
			}
		}
	}
	return pairs
}

func generateAntinodes(radarMap RadarMap, antennas map[string][]vec) map[utility.Point]int {
	antinodeSet := make(map[utility.Point]int)
	for key, positions := range antennas {

		if len(positions) < 2 {
			continue
		}

		pairs := generatePairs(positions)
		for _, pair := range pairs {
			a := pair[0]
			b := pair[1]
			diff := a.Sub(b) // Vector from b to a

			possibleAntinodes := []vec{b.Add(diff), b.Add(diff.Clone().Invert())}

			for _, antinode := range possibleAntinodes {

				if antinode.X() < 0 || antinode.X() >= float64(radarMap.width) {
					continue
				}
				if antinode.Y() < 0 || antinode.Y() >= float64(radarMap.height) {
					continue
				}
				if radarMap.data[int(antinode.Y())][int(antinode.X())] == key {
					continue
				}
				antinodeSet[utility.Point{X: int(antinode.X()), Y: int(antinode.Y())}]++
			}
		}
	}
	return antinodeSet
}

func printMap(radarMap RadarMap) {
	for _, row := range radarMap.data {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}

func Day8Part1() {

	radarMap := parseInput("day8/input.txt")
	printMap(radarMap)

	antennas := getAntennas(radarMap)
	antinodeMap := generateAntinodes(radarMap, antennas)

	numAntinodes := len(antinodeMap)

	fmt.Println(numAntinodes)

}
