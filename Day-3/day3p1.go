package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var area [1000][1000]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		var id, x, y, w, h int

		fmt.Sscanf(str, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				area[x+j][y+i]++
			}
		}
	}

	overlapArea := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if area[i][j] > 1 {
				overlapArea++
			}
		}
	}
	fmt.Println("Overlap area = ", overlapArea)
}
