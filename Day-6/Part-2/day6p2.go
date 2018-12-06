package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var points [50][2]int
	maxX, maxY := 0, 0
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		var x, y int
		fmt.Sscanf(input, "%d, %d", &x, &y)
		points[i][0] = x
		points[i][1] = y
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		i++
	}

	table := make([][]int, maxX+1)
	for i := range table {
		table[i] = make([]int, maxY+1)
	}

	regionCount := 0
	for i := range table {
		for j := range table[i] {
			totalDistant := 0
			for _, p := range points {
				distant := math.Abs(float64(i-p[0])) + math.Abs(float64(j-p[1]))
				totalDistant += int(distant)
			}
			if totalDistant < 10000 {
				regionCount++
			}
		}
	}

	fmt.Println("Answer =", regionCount)
}
