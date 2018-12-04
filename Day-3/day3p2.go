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

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		var id, x, y, w, h int
		isOverlap := false

		fmt.Sscanf(str, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		for i := 0; !isOverlap && i < h; i++ {
			for j := 0; !isOverlap && j < w; j++ {
				if area[x+j][y+i] > 1 {
					isOverlap = true
				}
			}
		}

		if !isOverlap {
			fmt.Println("Not overlap ID = ", id)
		}
	}
}
