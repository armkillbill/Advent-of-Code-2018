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

	for i := range table {
		for j := range table[i] {
			minDistant := math.MaxInt32
			var minIndex int
			moreThanOne := false
			for k, p := range points {
				distant := math.Abs(float64(i-p[0])) + math.Abs(float64(j-p[1]))
				if int(distant) < minDistant {
					minDistant = int(distant)
					minIndex = k
					moreThanOne = false
				} else if int(distant) == minDistant {
					moreThanOne = true
				}
			}
			if moreThanOne {
				table[i][j] = -1
			} else {
				table[i][j] = minIndex
			}
		}
	}
	// fmt.Println(table)

	maxCount := math.MinInt32
	for k := range points {
		count := 0
		for i := range table {
			for j := range table[i] {
				if table[i][j] == k {
					count++
				}
			}
		}
		if !isInfinite(k, table) && count > maxCount {
			maxCount = count
		}
	}

	fmt.Println("Answer =", maxCount)
}

func isInfinite(num int, table [][]int) bool {
	// Check top
	for _, v := range table[0] {
		if v == num {
			return true
		}
	}

	// Check left
	for i := range table {
		if table[i][0] == num {
			return true
		}
	}

	// Check right
	for i := range table {
		if table[i][len(table[i])-1] == num {
			return true
		}
	}

	// Check buttom
	for _, v := range table[len(table)-1] {
		if v == num {
			return true
		}
	}

	return false
}
