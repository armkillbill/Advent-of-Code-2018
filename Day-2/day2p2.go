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

	const maxList = 250
	const maxChar = 26
	var list [maxList]string
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); {
		str := scanner.Text()
		list[i] = str
		i++
	}

	maxMatch := -1
	var maxi, maxj int
	for i := 0; i < maxList-1; i++ {
		for j := i + 1; j < maxList; j++ {
			numMatch := 0
			for k := 0; k < maxChar; k++ {
				if list[i][k] == list[j][k] {
					numMatch++
				}
			}
			if numMatch > maxMatch {
				maxMatch = numMatch
				maxi = i
				maxj = j
			}
		}
	}

	for i := 0; i < maxChar; i++ {
		if list[maxi][i] == list[maxj][i] {
			fmt.Printf("%c", list[maxi][i])
		}
	}
}
